// Package topsecretsplit provee los dto de comunicación entre los satélites y el servicio topsecret_split
package topsecretsplit //import "github.com/aferrercrafter/operacion-fuego-quazar/internal/topsecret_split"
import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache exporta la instancia de cache
var Cache = cache.New(30*time.Minute, 30*time.Minute)

// SetCache setea el valor de una key
func SetCache(key string, sat Satellite) bool {
	Cache.Set(key, sat, cache.NoExpiration)
	return true
}

// GetCache obtiene el valor de una key
func GetCache(key string) (Satellite, bool) {
	var sat Satellite
	var found bool
	data, found := Cache.Get(key)
	if found {
		sat = data.(Satellite)
	}
	return sat, found
}
