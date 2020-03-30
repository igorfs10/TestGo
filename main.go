package main

import (
	"fmt"
	"math/rand"
	"time"

	inimigo "github.com/igorfs10/TestGo/structs"
)

// Inimigo : importação
type Inimigo = inimigo.Inimigo

const test string = "Teste"

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Teste", time.Now())

	var ola = Inimigo{0, "Rato", 5, 1, 1, 2, 0}
	ola.Nome = test

	fmt.Println("Teste", ola.Nome)
}
