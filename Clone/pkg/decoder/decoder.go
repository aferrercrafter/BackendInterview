// Package decoder provee funciones para obtener el mensaje escondido dentro
// de un arreglo de mensajes que tienen un defasaje y palabras irreconocibles
// que serán representadas como espacios en blanco
package decoder // import "github.com/aferrercrafter/operacion-fuego-quazar/pkg/decoder"

import (
	"fmt"
	"strings"
)

//GetMessage decodifica y retorna el mensaje escondido dentro de mensajes incompletos
// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (msg string) {

	if messages == nil {
		fmt.Println("err")
		return ""
	}

	_, min := getMaxMin(messages)

	//Defasaje 0 permitido
	//if max == min {
	//	return ""
	//}

	var decodedMsg []string
	var offset int = 0

	for i := 0; i < min; i++ {
		for _, msg := range messages {
			offset = len(msg) - min
			if len(msg[i+offset]) > 0 {
				decodedMsg = append(decodedMsg, msg[i+offset])
				break
			}
		}
	}

	return strings.Join(decodedMsg, " ")
}

func getMaxMin(messages [][]string) (int, int) {
	var min int = len(messages[0])
	var max int = 0
	for _, msg := range messages {
		if len(msg) < min {
			min = len(msg)
		}

		if len(msg) > max {
			max = len(msg)
		}
	}
	return max, min
}
