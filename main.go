package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/igorfs10/TestGo/consts"
	"github.com/igorfs10/TestGo/structs"
	"github.com/igorfs10/TestGo/utils"
)

// Inimigo : importação
type Inimigo = structs.Inimigo

// Personagem : importação
type Personagem = structs.Personagem

// Item : importação
type Item = structs.Item

func main() {
	fmt.Println("My favorite number is", rand.Intn(10), "na hora", time.Now())

	var outroInimigo Inimigo
	utils.JSONToVar(&outroInimigo, utils.CreatePath(1, consts.InimigoRato))
	fmt.Println(outroInimigo)

	var per Personagem
	utils.JSONToVar(&per, "data/personagens/guerreiro.json")
	fmt.Println(per)

	var it Item
	utils.JSONToVar(&it, "data/itens/pocao.json")
	fmt.Println(it)
}
