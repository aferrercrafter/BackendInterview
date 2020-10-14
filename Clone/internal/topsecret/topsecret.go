// Package topsecret provee los dto de comunicación entre los satélites y el servicio topsecret
package topsecret

import (
	"strings"

	"github.com/aferrercrafter/operacion-fuego-quazar/pkg/decoder"
	"github.com/aferrercrafter/operacion-fuego-quazar/pkg/locator"
)

// GetTopsecret obtiene la posicion y mensaje ocultos dado un request valido
// input: request con la distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetTopsecret(satellites SatellitesRequest) (x, y float32, message string) {

	var r1, r2, r3 float32
	var messages = make(map[string][]string)

	for _, satellite := range satellites.Satellites {

		switch strings.ToLower(satellite.Name) {
		case "kenobi":
			r1 = satellite.Distance
			messages["kenobi"] = satellite.Message
		case "skywalker":
			r2 = satellite.Distance
			messages["skywalker"] = satellite.Message
		case "sato":
			r3 = satellite.Distance
			messages["sato"] = satellite.Message
		}
	}

	resX, resY := locator.GetLocation(r1, r2, r3)
	msg := decoder.GetMessage(messages["kenobi"], messages["skywalker"], messages["sato"])

	return resX, resY, msg
}
