package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generar reporte a partir de datos de escaneo",
	Long:  "Lee un archivo JSON de resultados y genera un reporte HTML.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[*] Funcionalidad de reporte — use scan --format json --output results.json")
		fmt.Println("[*] Luego convierta con: netprobe report --input results.json --format html")
	},
}
