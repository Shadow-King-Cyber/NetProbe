# NetProbe

Scanner de red de alto rendimiento escrito en Go — escaneo concurrente de puertos con goroutines.

> **ADVERTENCIA**: Solo para redes autorizadas. El escaneo no autorizado es ilegal.

## Características

- **Escaneo concurrente** con goroutines (miles de puertos en paralelo)
- **Service fingerprinting** — detección de servicios por puerto conocido
- **Timeout configurable** por puerto
- **Exportación** a JSON y HTML con estilos dark theme
- **CLI con Cobra** — comandos intuitivos
- **Rendimiento** — binario estático sin dependencias externas

## Aviso Legal

Esta herramienta se proporciona únicamente con fines educativos y para pruebas de seguridad autorizadas.

**Al usar este software, aceptas que:**
- Solo lo usarás en redes que poseas o para las que tengas autorización explícita
- El escaneo no autorizado de redes es ilegal en la mayoría de jurisdicciones
- Los autores no asumen responsabilidad por uso indebido

## Requisitos

- Go 1.21+

```bash
git clone https://github.com/Shadow-King-Cyber/NetProbe.git
cd NetProbe
go build -o netprobe .
```

## Inicio Rápido

```bash
# Escanear puertos comunes (1-1024)
./netprobe scan --host 127.0.0.1

# Escanear rango específico con más workers
./netprobe scan --host 192.168.1.1 --start-port 1 --end-port 65535 --workers 500

# Exportar resultados a JSON
./netprobe scan --host 10.0.0.1 --format json --output resultados.json

# Exportar a HTML
./netprobe scan --host 10.0.0.1 --format html --output reporte.html
```

## Comandos del CLI

```bash
# Escaneo básico
./netprobe scan --host <IP>

# Escaneo completo con opciones
./netprobe scan --host <IP> --start-port 1 --end-port 65535 --workers 500 --timeout 1000

# Exportar resultados
./netprobe scan --host <IP> --format json --output scan.json
./netprobe scan --host <IP> --format html --output scan.html

# Versión
./netprobe version
```

## Estructura del Proyecto

```
NetProbe/
├── main.go              # Punto de entrada
├── cmd/
│   ├── root.go          # Comando raíz (Cobra)
│   ├── scan.go          # Comando de escaneo
│   └── report.go        # Comando de reportes
├── scanner/
│   ├── scanner.go       # Motor de escaneo TCP concurrente
│   └── scanner_test.go  # Tests
├── reporter/
│   └── reporter.go      # Exportación JSON/HTML
├── go.mod               # Dependencias Go
├── .gitignore
├── LICENSE              # Licencia MIT
└── README.md
```

## Ejecutar Tests

```bash
go test ./...
```

## Licencia

MIT License — ver [LICENSE](LICENSE)
