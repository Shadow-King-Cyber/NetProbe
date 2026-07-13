package scanner

import (
	"testing"
	"time"
)

func TestObtenerServicioHTTP(t *testing.T) {
	servicio := ObtenerServicio(80)
	if servicio != "HTTP" {
		t.Errorf("Se esperaba HTTP, se obtuvo %s", servicio)
	}
}

func TestObtenerServicioSSH(t *testing.T) {
	servicio := ObtenerServicio(22)
	if servicio != "SSH" {
		t.Errorf("Se esperaba SSH, se obtuvo %s", servicio)
	}
}

func TestObtenerServicioDesconocido(t *testing.T) {
	servicio := ObtenerServicio(99999)
	if servicio != "desconocido" {
		t.Errorf("Se esperaba desconocido, se obtuvo %s", servicio)
	}
}

func TestObtenerServicioHTTPS(t *testing.T) {
	servicio := ObtenerServicio(443)
	if servicio != "HTTPS" {
		t.Errorf("Se esperaba HTTPS, se obtuvo %s", servicio)
	}
}

func TestEscanearPuertoCerrado(t *testing.T) {
	resultado := EscanearPuerto("127.0.0.1", 1, 100*time.Millisecond)
	if resultado.Abierto {
		t.Error("Se esperaba que el puerto 1 estuviera cerrado")
	}
	if resultado.Estado != "cerrado" {
		t.Errorf("Se esperaba estado 'cerrado', se obtuvo %s", resultado.Estado)
	}
}

func TestEscanearPuertoHostInvalido(t *testing.T) {
	resultado := EscanearPuerto("192.0.2.1", 80, 100*time.Millisecond)
	if resultado.Abierto {
		t.Error("Se esperaba que un host inválido tuviera el puerto cerrado")
	}
}
