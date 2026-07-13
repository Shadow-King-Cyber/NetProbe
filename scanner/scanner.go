// Package scanner — Motor de escaneo de puertos TCP concurrente.
package scanner

import (
	"fmt"
	"net"
	"time"
)

// ResultadoPuerto representa el resultado del escaneo de un puerto individual.
type ResultadoPuerto struct {
	Host    string `json:"host"`
	Puerto  int    `json:"puerto"`
	Abierto bool   `json:"abierto"`
	Estado  string `json:"estado"`
	Servicio string `json:"servicio"`
	Tiempo  time.Duration `json:"tiempo"`
}

// servicios conocidos — mapeo de puerto a nombre de servicio.
var serviciosConocidos = map[int]string{
	21:    "FTP",
	22:    "SSH",
	23:    "Telnet",
	25:    "SMTP",
	53:    "DNS",
	80:    "HTTP",
	110:   "POP3",
	143:   "IMAP",
	443:   "HTTPS",
	993:   "IMAPS",
	995:   "POP3S",
	3306:  "MySQL",
	3389:  "RDP",
	5432:  "PostgreSQL",
	6379:  "Redis",
	8080:  "HTTP-Alt",
	8443:  "HTTPS-Alt",
	27017: "MongoDB",
}

// EscanearPuerto intenta conectarse a un puerto TCP específico.
// Devuelve un ResultadoPuerto con el estado del puerto.
func EscanearPuerto(host string, puerto int, timeout time.Duration) ResultadoPuerto {
	addr := fmt.Sprintf("%s:%d", host, puerto)
	inicio := time.Now()

	conn, err := net.DialTimeout("tcp", addr, timeout)
	tiempo := time.Since(inicio)

	resultado := ResultadoPuerto{
		Host:   host,
		Puerto: puerto,
		Tiempo: tiempo,
	}

	if err != nil {
		resultado.Abierto = false
		resultado.Estado = "cerrado"
		return resultado
	}

	conn.Close()

	resultado.Abierto = true
	resultado.Estado = "abierto"
	resultado.Servicio = ObtenerServicio(puerto)

	return resultado
}

// ObtenerServicio devuelve el nombre del servicio asociado a un puerto.
func ObtenerServicio(puerto int) string {
	if servicio, ok := serviciosConocidos[puerto]; ok {
		return servicio
	}
	return "desconocido"
}
