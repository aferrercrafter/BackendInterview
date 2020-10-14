// Package locator provee funciones para obtener la posición de un punto desconocido
// en un eje de coordenadas (x, y), teniendo como información la distancia de dicho punto
// con al menos tres puntos conocidos
package locator // import "github.com/aferrercrafter/operacion-fuego-quazar/pkg/locator"
import (
	"fmt"
	"strconv"
)

// GetLocation obtiene la posicion dada la distancia hacia los 3 satélites
// usando el método de trilateración bidimensional
// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float32) (x, y float32) {

	if distances == nil {
		return 0.0, 0.0
	}

	if len(distances) < 3 {
		return 0.0, 0.0
	}

	var (
		x1 float32 = Kenobi.X
		y1 float32 = Kenobi.Y

		x2 float32 = Skywalker.X
		y2 float32 = Skywalker.Y

		x3 float32 = Sato.X
		y3 float32 = Sato.Y

		r1 float32 = distances[0]
		r2 float32 = distances[1]
		r3 float32 = distances[2]

		x1Sq float32 = x1 * x1
		x2Sq float32 = x2 * x2
		x3Sq float32 = x3 * x3
		y1Sq float32 = y1 * y1
		y2Sq float32 = y2 * y2
		y3Sq float32 = y3 * y3
		r1Sq float32 = r1 * r1
		r2Sq float32 = r2 * r2
		r3Sq float32 = r3 * r3
	)

	var num1 float32 = (x2-x1)*(x3Sq+y3Sq-r3Sq) + (x1-x3)*(x2Sq+y2Sq-r2Sq) + (x3-x2)*(x1Sq+y1Sq-r1Sq)
	var den1 float32 = 2 * (y3*(x2-x1) + y2*(x1-x3) + y1*(x3-x2))
	var resY = num1 / den1
	var num2 float32 = r2Sq - r1Sq + x1Sq - x2Sq + y1Sq - y2Sq - 2*(y1-y2)*resY
	var den2 float32 = 2 * (x1 - x2)
	var resX = num2 / den2

	var resYF string = fmt.Sprintf("%.1f", resY)
	var resXF string = fmt.Sprintf("%.1f", resX)

	var resY64 float64
	var resX64 float64

	var err error

	resY64, err = strconv.ParseFloat(resYF, 32)
	if err != nil {
		return 0.0, 0.0
	}
	resX64, err = strconv.ParseFloat(resXF, 32)
	if err != nil {
		return 0.0, 0.0
	}

	return float32(resX64), float32(resY64)
}
