// main package para el servidor Topsecret.
package main

import "github.com/aferrercrafter/operacion-fuego-quazar/api/server"

// @title fuego-quazar API
// @version 1.0
// @description Servicio topsecret para operación fuego quazar.
// @termsOfService http://swagger.io/terms/

// @contact.name Arnaldo Corzo
// @contact.url https://github.com/aferrercrafter
// @contact.email aferrercrafter@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

func main() {
	server.Start()
}
