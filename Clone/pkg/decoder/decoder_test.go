// Package decoder provee funciones para obtener el mensaje escondido dentro
// de un arreglo de mensajes que tienen un defasaje y palabras irreconocibles
// que serán representadas como espacios en blanco
package decoder // import "github.com/aferrercrafter/operacion-fuego-quazar/pkg/decoder"

import (
	"testing"
)

func TestGetMessage(t *testing.T) {

	// Caso exitoso
	decodedMsg := GetMessage([]string{"", "este", "es", "un", "mensaje"}, []string{"este", "", "un", "mensaje"}, []string{"", "", "es", "", "mensaje"})
	if decodedMsg != "este es un mensaje" {
		t.Fatalf("Fail no se decodificó correctamente el mensaje")
	}

	// Caso argumentos nil
	nilMsg := GetMessage()
	if nilMsg != "" {
		t.Fatalf("Fail debe retornar vacío al no tener mensajes válidos (mensajes nil)")
	}

}
