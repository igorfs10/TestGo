package main

import (
	"fmt"
	"math/rand"
	"time"

	inimigo "github.com/igorfs10/TestGo/structs"
	file "github.com/igorfs10/TestGo/utils"
)

// Inimigo : importação
type Inimigo = inimigo.Inimigo

const test string = "Teste"

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Teste", time.Now())

	var ola Inimigo
	ola.Nome = test

	fmt.Println("Teste", ola.Nome)

	var outroInimigo Inimigo
	file.JSONToVar(&outroInimigo)
	fmt.Println(outroInimigo)
}
