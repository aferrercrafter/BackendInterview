// Package topsecretsplit provee los dto de comunicación entre los satélites y el servicio topsecret_split
package topsecretsplit //import "github.com/aferrercrafter/operacion-fuego-quazar/internal/topsecret_split"
import (
	"fmt"

	"github.com/aferrercrafter/operacion-fuego-quazar/pkg/decoder"
	"github.com/aferrercrafter/operacion-fuego-quazar/pkg/locator"
)

// PostTopSecretSplit Crea la llave de uno de los satellites válidos aun no creado
func PostTopSecretSplit(name string, sat Satellite) error {

	_, found := GetCache(name)

	if found {
		return fmt.Errorf("satélite %s ya está creado", name)
	}

	SetCache(name, sat)

	return nil
}

// GetTopSecretSplit obtiene la posicion y mensaje ocultos de los satelites guardados en cache
// input: request con la distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetTopSecretSplit() (x, y float32, message string, err error) {

	var kenobi, foundKenobi = GetCache("kenobi")
	var skywalker, foundSkywalker = GetCache("skywalker")
	var sato, foundSato = GetCache("sato")

	if !foundKenobi {
		return 0, 0, "", fmt.Errorf("No hay suficiente información Kenobi missing")
	}

	if !foundSkywalker {
		return 0, 0, "", fmt.Errorf("No hay suficiente información Skywalker missing")
	}

	if !foundSato {
		return 0, 0, "", fmt.Errorf("No hay suficiente información Sato missing")
	}

	resX, resY := locator.GetLocation(kenobi.Distance, skywalker.Distance, sato.Distance)
	msg := decoder.GetMessage(kenobi.Message, skywalker.Message, sato.Message)

	return resX, resY, msg, nil

}
