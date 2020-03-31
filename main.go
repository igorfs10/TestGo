package main

import (
	"fmt"
	"math/rand"
	"time"

	structs "github.com/igorfs10/TestGo/structs"
	file "github.com/igorfs10/TestGo/utils"
)

// Inimigo : importação
type Inimigo = structs.Inimigo

// Personagem : importação
type Personagem = structs.Personagem

const test string = "Teste"

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Teste", time.Now())

	var ola Inimigo
	ola.Nome = test

	fmt.Println("Teste", ola.Nome)

	var outroInimigo Inimigo
	file.JSONToVar(&outroInimigo, "data/inimigos/rato.json")
	fmt.Println(outroInimigo)

	var per Personagem
	file.JSONToVar(&per, "data/personagens/guerreiro.json")

	fmt.Println(per)
}
