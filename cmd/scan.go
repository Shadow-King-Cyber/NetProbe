package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Shadow-King-Cyber/NetProbe/reporter"
	"github.com/Shadow-King-Cyber/NetProbe/scanner"
	"github.com/spf13/cobra"
)

var (
	targetHost string
	startPort  int
	endPort    int
	workers    int
	timeout    int
	outputFmt  string
	outputFile string
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Escanear puertos en un host objetivo",
	Long:  "Escanea puertos TCP en un host usando concurrencia con goroutines.",
	RunE:  runScan,
}

func init() {
	scanCmd.Flags().StringVarP(&targetHost, "host", "H", "", "Host objetivo (IP o dominio)")
	scanCmd.Flags().IntVarP(&startPort, "start-port", "s", 1, "Puerto inicial")
	scanCmd.Flags().IntVarP(&endPort, "end-port", "e", 1024, "Puerto final")
	scanCmd.Flags().IntVarP(&workers, "workers", "w", 100, "Número de goroutines concurrentes")
	scanCmd.Flags().IntVarP(&timeout, "timeout", "t", 2000, "Timeout por puerto en milisegundos")
	scanCmd.Flags().StringVarP(&outputFmt, "format", "f", "text", "Formato de salida: text, json, html")
	scanCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Archivo de salida (vacío = stdout)")
	_ = scanCmd.MarkFlagRequired("host")
}

func runScan(cmd *cobra.Command, args []string) error {
	fmt.Printf("[*] Escaneando %s:%d-%d con %d workers...\n", targetHost, startPort, endPort, workers)

	inicio := time.Now()

	// Crear canal de puertos y canal de resultados
	puertos := make(chan int, workers)
	resultados := make(chan scanner.ResultadoPuerto, endPort-startPort+1)

	// Lanzar workers
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for puerto := range puertos {
				resultado := scanner.EscanearPuerto(targetHost, puerto, time.Duration(timeout)*time.Millisecond)
				resultados <- resultado
			}
		}()
	}

	// Enviar puertos a escanear
	go func() {
		for p := startPort; p <= endPort; p++ {
			puertos <- p
		}
		close(puertos)
	}()

	// Esperar a que terminen los workers y cerrar resultados
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Recopilar resultados
	var listaResultados []scanner.ResultadoPuerto
	for r := range resultados {
		if r.Abierto {
			listaResultados = append(listaResultados, r)
		}
	}

	duracion := time.Since(inicio)

	fmt.Printf("[+] Escaneo completado en %v\n", duracion)
	fmt.Printf("[+] Puertos abiertos: %d\n\n", len(listaResultados))

	// Mostrar resultados
	for _, r := range listaResultados {
		fmt.Printf("  %d/tcp  %-12s  %s\n", r.Puerto, r.Estado, r.Servicio)
	}

	// Exportar si se solicita
	if outputFile != "" {
		err := reporter.Exportar(listaResultados, outputFile, outputFmt, duracion)
		if err != nil {
			return fmt.Errorf("error exportando: %w", err)
		}
		fmt.Printf("\n[+] Reporte guardado en: %s", outputFile)
	}

	return nil
}

// parsearRangoPuertos convierte un string como "80,443,8000-9000" a una lista de enteros.
func parsearRangoPuertos(rango string) []int {
	var puertos []int
	for _, parte := range strings.Split(rango, ",") {
		parte = strings.TrimSpace(parte)
		if strings.Contains(parte, "-") {
			limits := strings.Split(parte, "-")
			if len(limits) == 2 {
				ini, _ := strconv.Atoi(limits[0])
				fin, _ := strconv.Atoi(limits[1])
				for p := ini; p <= fin; p++ {
					puertos = append(puertos, p)
				}
			}
		} else {
			p, _ := strconv.Atoi(parte)
			puertos = append(puertos, p)
		}
	}
	return puertos
}
