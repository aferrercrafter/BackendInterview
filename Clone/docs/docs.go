// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Arnaldo Corzo",
            "url": "https://github.com/aferrercrafter",
            "email": "aferrercrafter@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/topsecret": {
            "post": {
                "description": "Obtiene posición y mensaje escondido",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene posición y mensaje escondido",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/topsecret.SatellitesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/topsecret_split": {
            "get": {
                "description": "Obtiene posición y mensaje escondido",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene posición y mensaje escondido",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/topsecretsplit.SatelliteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Añade la información de un satélite",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Añade la información de un satélite",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/topsecretsplit.Satellite"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "topsecret.Position": {
            "type": "object",
            "required": [
                "x",
                "y"
            ],
            "properties": {
                "x": {
                    "type": "string"
                },
                "y": {
                    "type": "string"
                }
            }
        },
        "topsecret.SatellitesResponse": {
            "type": "object",
            "required": [
                "message",
                "position"
            ],
            "properties": {
                "message": {
                    "type": "string"
                },
                "position": {
                    "type": "object",
                    "$ref": "#/definitions/topsecret.Position"
                }
            }
        },
        "topsecretsplit.Position": {
            "type": "object",
            "required": [
                "x",
                "y"
            ],
            "properties": {
                "x": {
                    "type": "string"
                },
                "y": {
                    "type": "string"
                }
            }
        },
        "topsecretsplit.Satellite": {
            "type": "object",
            "required": [
                "distance",
                "message"
            ],
            "properties": {
                "distance": {
                    "type": "number"
                },
                "message": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "topsecretsplit.SatelliteResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "position": {
                    "type": "object",
                    "$ref": "#/definitions/topsecretsplit.Position"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "fuego-quazar API",
	Description: "Servicio topsecret para operación fuego quazar.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}