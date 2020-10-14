// Package topsecret provee los dto de comunicación entre los satélites y el servicio topsecret
package topsecret //import "github.com/aferrercrafter/operacion-fuego-quazar/internal/topsecret"
import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// Satellites dto
type Satellites struct {
	Name     string   `json:"name" validate:"required"`
	Distance float32  `json:"distance" validate:"required"`
	Message  []string `json:"message" validate:"required"`
}

// SatellitesRequest dto
type SatellitesRequest struct {
	Satellites []Satellites `json:"satellites" validate:"gt=0,dive,len=1,dive,required"`
}

// Position dto
type Position struct {
	X json.Number `json:"x" validate:"required"`
	Y json.Number `json:"y" validate:"required"`
}

// SatellitesResponse dto
type SatellitesResponse struct {
	Position Position `json:"position" validate:"required"`
	Message  string   `json:"message" validate:"required"`
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
