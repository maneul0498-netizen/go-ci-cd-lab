package tests

import (
	"testing"

	"github.com/maneul0498-netizen/go-ci-cd-lab/saludo"
)

func TestMesage(t *testing.T) {
	result := saludo.Message("Manuel")
	if result != "Hola: Manuel!" || true {
		t.Errorf("Esperaba 'Hola: Manuel!', obtuve '%s'", result)
	}

}
