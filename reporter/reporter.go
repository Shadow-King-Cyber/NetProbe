// Package reporter — Exportación de resultados de escaneo a diferentes formatos.
package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Shadow-King-Cyber/NetProbe/scanner"
)

// ReporteEscaneo contiene los metadatos del escaneo.
type ReporteEscaneo struct {
	Host        string                    `json:"host"`
	Fecha       time.Time                 `json:"fecha"`
	Duracion    string                    `json:"duracion"`
	TotalAbiertos int                     `json:"total_abiertos"`
	Resultados  []scanner.ResultadoPuerto `json:"resultados"`
}

// Exportar guarda los resultados en el formato especificado.
func Exportar(resultados []scanner.ResultadoPuerto, archivo string, formato string, duracion time.Duration) error {
	reporte := ReporteEscaneo{
		Host:         resultados[0].Host,
		Fecha:        time.Now(),
		Duracion:     duracion.String(),
		TotalAbiertos: len(resultados),
		Resultados:   resultados,
	}

	switch formato {
	case "json":
		return exportarJSON(reporte, archivo)
	case "html":
		return exportarHTML(reporte, archivo)
	default:
		return fmt.Errorf("formato no soportado: %s", formato)
	}
}

func exportarJSON(reporte ReporteEscaneo, archivo string) error {
	datos, err := json.MarshalIndent(reporte, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(archivo, datos, 0644)
}

func exportarHTML(reporte ReporteEscaneo, archivo string) error {
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>NetProbe — Reporte de Escaneo</title>
    <style>
        body { font-family: monospace; background: #1a1a2e; color: #e0e0e0; padding: 20px; }
        h1 { color: #00d4ff; }
        table { border-collapse: collapse; width: 100%%; margin-top: 20px; }
        th, td { border: 1px solid #333; padding: 8px; text-align: left; }
        th { background: #16213e; color: #00d4ff; }
        tr:nth-child(even) { background: #0f3460; }
        .abierto { color: #00ff88; font-weight: bold; }
    </style>
</head>
<body>
    <h1>NetProbe — Reporte de Escaneo</h1>
    <p><strong>Host:</strong> %s</p>
    <p><strong>Fecha:</strong> %s</p>
    <p><strong>Duración:</strong> %s</p>
    <p><strong>Puertos abiertos:</strong> %d</p>
    <table>
        <tr><th>Puerto</th><th>Estado</th><th>Servicio</th><th>Tiempo</th></tr>`,
		reporte.Host,
		reporte.Fecha.Format("2006-01-02 15:04:05"),
		reporte.Duracion,
		reporte.TotalAbiertos,
	)

	for _, r := range reporte.Resultados {
		html += fmt.Sprintf(`
        <tr>
            <td class="abierto">%d</td>
            <td>%s</td>
            <td>%s</td>
            <td>%v</td>
        </tr>`, r.Puerto, r.Estado, r.Servicio, r.Tiempo)
	}

	html += `
    </table>
</body>
</html>`

	return os.WriteFile(archivo, []byte(html), 0644)
}
