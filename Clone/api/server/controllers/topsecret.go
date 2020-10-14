// Package controllers provee los handlers de cada controlador
package controllers // import "github.com/aferrercrafter/operacion-fuego-quazar/api/controllers"

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	httputil "github.com/aferrercrafter/operacion-fuego-quazar/api/httputils"
	"github.com/aferrercrafter/operacion-fuego-quazar/internal/topsecret"
	"github.com/gin-gonic/gin"
)

// TopSecretPost godoc
// @Summary Obtiene posición y mensaje escondido
// @Description Obtiene posición y mensaje escondido
// @Accept  json
// @Produce  json
// @Success 200 {object} topsecret.SatellitesResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /topsecret [post]
func TopSecretPost(c *gin.Context) {

	var satelliteRequest topsecret.SatellitesRequest

	if err := c.ShouldBindJSON(&satelliteRequest); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if len(satelliteRequest.Satellites) < 3 {
		httputil.NewError(c, http.StatusBadRequest, errors.New("no hay suficientes satélites"))
		return
	}

	if err := topsecret.Validate.Struct(satelliteRequest); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	x, y, msg := topsecret.GetTopsecret(satelliteRequest)

	var resp topsecret.Position
	resp.X = json.Number(fmt.Sprintf("%.1f", x))
	resp.Y = json.Number(fmt.Sprintf("%.1f", y))

	c.JSON(http.StatusOK, gin.H{"position": resp, "message": msg})
}
