// Package locator provee funciones para obtener la posición de un punto desconocido
// en un eje de coordenadas (x, y), teniendo como información la distancia de dicho punto
// con al menos tres puntos conocidos
package locator // import "github.com/aferrercrafter/operacion-fuego-quazar/pkg/locator"

import (
	"fmt"
	"testing"
)

func TestGetLocation(t *testing.T) {

	// Caso exitoso
	x, y := GetLocation(100.0, 115.5, 142.7)
	if fmt.Sprintf("%.1f", x) != "-487.3" || fmt.Sprintf("%.1f", y) != "1557.0" {
		t.Fatalf("Fail no se encontraron las coordenadas correctas")
	}

	// Caso argumentos nil
	nilX, nilY := GetLocation()
	if nilX != 0 || nilY != 0 {
		t.Fatalf("Fail debe retornar coordenadas (0, 0) por no tener distancias válidas (nil)")
	}

	// Caso sin defasaje
	fewX, fewY := GetLocation(100.0, 115.5)
	if fewX != 0 || fewY != 0 {
		t.Fatalf("Fail debe retornar coordenadas (0, 0) por no tener distancias válidas (solo 2 distancias)")
	}

}
