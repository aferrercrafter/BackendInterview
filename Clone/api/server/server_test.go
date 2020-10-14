package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()

	var jsonStr []byte = []byte("{\"satellites\": [{\"name\": \"kenobi\",\"distance\": 100.0,\"message\": [\"este\", \"\", \"\", \"mensaje\", \"\"]},{\"name\": \"skywalker\",\"distance\": 115.5,\"message\": [\"\", \"es\", \"\", \"\", \"secreto\"]},{\"name\": \"sato\",\"distance\": 142.7,\"message\": [\"este\", \"\", \"un\", \"\", \"\"]}]}")

	req, _ := http.NewRequest("POST", "/api/v1/topsecret", bytes.NewBuffer(jsonStr))

	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Fail debe retornar status code 200")
	}

	if w.Body.String() != "{\"message\":\"este es un mensaje secreto\",\"position\":{\"x\":-487.3,\"y\":1557.0}}" {
		t.Fatalf("Fail debe respuesta valida")
	}

	req1, _ := http.NewRequest("GET", "/api/v1/topsecret_split", bytes.NewBuffer(jsonStr))

	w = httptest.NewRecorder()

	router.ServeHTTP(w, req1)

	if w.Code != 400 {
		t.Fatalf("Fail debe retornar status code 400 por no tener satellites")
	}

}
