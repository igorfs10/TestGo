package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/igorfs10/TestGo/structs"
)

const pastaData string = "data/"
const pastaPersonagem string = "personagens/"
const pastaInimigo string = "inimigos/"
const pastaItem string = "itens/"

// TipoPersonagem : id do tipo personagem
const TipoPersonagem int8 = 0

// TipoInimigo : id do tipo inimigo
const TipoInimigo int8 = 1

// TipoItem : id do tipo item
const TipoItem int8 = 2

// JSONParaVar : Carrega um arquivo JSON e coloca numa variável. Ele usa o ponteiro da variável e o caminho do arquivo JSON.
func JSONParaVar(variavel interface{}, caminho string) {
	jsonFile, _ := os.Open(caminho)
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &variavel)
}

// CriarCaminho : Cria um caminho passando o tipo de objeto que quer pegar e o id do objeto
func CriarCaminho(tipo int8, id int8) string {
	var caminho string = pastaData

	switch tipo {
	/*-----		Personagem	-----*/
	case TipoPersonagem:
		caminho += pastaPersonagem
		switch id {
		case structs.PersonagemGuerreiro:
			caminho += "guerreiro"
		}
	/*-----		Inimigo		-----*/
	case TipoInimigo:
		caminho += pastaInimigo
		switch id {
		case structs.InimigoRato:
			caminho += "rato"
		case structs.InimigoCoelho:
			caminho += "coelho"
		case structs.InimigoLobo:
			caminho += "lobo"
		}

	/*-----		Item		-----*/
	case TipoItem:
		caminho += pastaItem
		switch id {
		case structs.ItemPocao:
			caminho += "pocao"
		}
	}

	caminho += ".json"

	return caminho
}
