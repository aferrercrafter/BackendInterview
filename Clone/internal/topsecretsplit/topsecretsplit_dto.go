// Package topsecretsplit provee los dto de comunicación entre los satélites y el servicio topsecret_split
package topsecretsplit //import "github.com/aferrercrafter/operacion-fuego-quazar/internal/topsecret_split"
import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// Satellite dto
type Satellite struct {
	Distance float32  `json:"distance" validate:"required"`
	Message  []string `json:"message" validate:"required"`
}

// Position dto
type Position struct {
	X json.Number `json:"x" validate:"required"`
	Y json.Number `json:"y" validate:"required"`
}

// SatelliteResponse dto
type SatelliteResponse struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

// Validate usado para validar la estructura del payload del request
var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("json", isJSON)
}

// isJSON chequea que sea un json válido
func isJSON(fl validator.FieldLevel) bool {
	var j map[string]interface{}
	if err := json.Unmarshal([]byte(fl.Field().String()), &j); err != nil {
		return false
	}
	return true
}
