// Package topsecretsplit provee los dto de comunicación entre los satélites y el servicio topsecret_split
package topsecretsplit

import (
	"fmt"
	"testing"
)

func TestPostTopSecretSplit(t *testing.T) {

	sat := Satellite{
		Distance: 100.0,
		Message:  []string{"este", "", "", "mensaje", ""},
	}

	err := PostTopSecretSplit("kenobi", sat)
	if err != nil {
		t.Fatalf("Fail nombre incorrecto de satelite")
	}
	_, found := GetCache("kenobi")
	if !found {
		t.Fatalf("Fail kenobi debería estar presente")
	}

}

func TestGetTopSecretSplit(t *testing.T) {

	sat2 := Satellite{
		Distance: 115.5,
		Message:  []string{"", "es", "", "", "secreto"},
	}
	sat3 := Satellite{
		Distance: 142.7,
		Message:  []string{"este", "", "un", "", ""},
	}

	_, _, _, err := GetTopSecretSplit()
	if err == nil {
		t.Fatalf("Fail debería ser información insuficiente")
	}

	_ = PostTopSecretSplit("skywalker", sat2)
	_ = PostTopSecretSplit("sato", sat3)

	x, y, msg, _ := GetTopSecretSplit()

	if fmt.Sprintf("%.1f", x) != "-487.3" {
		t.Fatalf("Fail valor X incorrecto")
	}

	if fmt.Sprintf("%.1f", y) != "1557.0" {
		t.Fatalf("Fail valor Y incorrecto")
	}

	if msg != "este es un mensaje secreto" {
		t.Fatalf("Fail valor msg incorrecto")
	}

}
