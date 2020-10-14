// Package topsecret provee los dto de comunicación entre los satélites y el servicio topsecret
package topsecret

import (
	"fmt"
	"testing"
)

func TestGetTopsecret(t *testing.T) {

	// Caso exitoso
	req := SatellitesRequest{
		[]Satellites{

			Satellites{
				Name:     "kenobi",
				Distance: 100.0,
				Message:  []string{"este", "", "", "mensaje", ""},
			},
			Satellites{
				Name:     "skywalker",
				Distance: 115.5,
				Message:  []string{"", "es", "", "", "secreto"},
			},
			Satellites{
				Name:     "sato",
				Distance: 142.7,
				Message:  []string{"este", "", "un", "", ""},
			},
		},
	}

	// Input válido
	X, Y, msg := GetTopsecret(req)

	if fmt.Sprintf("%.1f", X) != "-487.3" {
		t.Fatalf("Fail resultado incorrecto de X")
	}
	if fmt.Sprintf("%.1f", Y) != "1557.0" {
		t.Fatalf("Fail resultado incorrecto de Y")
	}
	if msg != "este es un mensaje secreto" {
		t.Fatalf("Fail resultado incorrecto de msg")
	}

}
