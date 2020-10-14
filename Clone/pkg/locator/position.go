// Package locator provee funciones para obtener la posición de un punto desconocido
// en un eje de coordenadas (x, y), teniendo como información la distancia de dicho punto
// con al menos tres puntos conocidos
package locator

// Position modelo de una coordenada
type Position struct {
	X float32
	Y float32
}

var (
	// Kenobi coordenadas
	Kenobi = Position{-500, -200}

	// Skywalker coordenadas
	Skywalker = Position{100, -100}

	// Sato coordenadas
	Sato = Position{500, 100}
)
