package main

import (
	"fmt"
	"math/rand"
	"time"

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
	utils.JSONParaVar(&outroInimigo, utils.CriarCaminho(utils.TipoInimigo, structs.InimigoRato))
	fmt.Println(outroInimigo)

	var per Personagem
	utils.JSONParaVar(&per, utils.CriarCaminho(utils.TipoPersonagem, structs.PersonagemGuerreiro))
	fmt.Println(per)

	var it Item
	utils.JSONParaVar(&it, utils.CriarCaminho(utils.TipoItem, structs.ItemPocao))
	fmt.Println(it)
}
