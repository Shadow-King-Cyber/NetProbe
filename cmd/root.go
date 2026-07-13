// Package cmd — Comandos CLI para NetProbe usando Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netprobe",
	Short: "NetProbe — Scanner de red de alto rendimiento",
	Long:  "NetProbe es un scanner de red concurrente escrito en Go. Escanea puertos, detecta servicios y genera reportes.",
}

// Execute ejecuta el comando raíz.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Mostrar versión de NetProbe",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NetProbe v1.0.0")
		fmt.Println("Go runtime:", fmt.Sprintf("%s", os.Getenv("GO_VERSION")))
	},
}
