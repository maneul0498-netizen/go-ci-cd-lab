package saludo

import "fmt"

func Message(name string) string {
	return fmt.Sprintf("Hola: %v!", name)
}
