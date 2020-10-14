// Package controllers provee los handlers de cada controlador
package controllers // import "github.com/aferrercrafter/operacion-fuego-quazar/api/controllers"

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	httputil "github.com/aferrercrafter/operacion-fuego-quazar/api/httputils"
	"github.com/aferrercrafter/operacion-fuego-quazar/internal/topsecret"
	"github.com/aferrercrafter/operacion-fuego-quazar/internal/topsecretsplit"
	"github.com/gin-gonic/gin"
)

// TopSecretSplitGet godoc
// @Summary Obtiene posición y mensaje escondido
// @Description Obtiene posición y mensaje escondido
// @Accept  json
// @Produce  json
// @Success 200 {object} topsecretsplit.SatelliteResponse
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /topsecret_split [get]
func TopSecretSplitGet(c *gin.Context) {

	x, y, msg, err := topsecretsplit.GetTopSecretSplit()

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	var resp topsecretsplit.Position
	resp.X = json.Number(fmt.Sprintf("%.1f", x))
	resp.Y = json.Number(fmt.Sprintf("%.1f", y))

	c.JSON(http.StatusOK, gin.H{"position": resp, "message": msg})

}

// TopSecretSplitPost godoc
// @Summary Añade la información de un satélite
// @Description Añade la información de un satélite
// @Accept  json
// @Produce  json
// @Success 200 {object} topsecretsplit.Satellite
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /topsecret_split [post]
func TopSecretSplitPost(c *gin.Context) {

	var name string = c.Param("name")
	var satelliteRequest topsecretsplit.Satellite

	if !(name == "kenobi" || name == "skywalker" || name == "sato") {
		httputil.NewError(c, http.StatusNotFound, errors.New("satélite inválido"))
		return
	}

	if err := c.ShouldBindJSON(&satelliteRequest); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if err := topsecret.Validate.Struct(satelliteRequest); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	err := topsecretsplit.PostTopSecretSplit(name, satelliteRequest)

	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, satelliteRequest)
}
