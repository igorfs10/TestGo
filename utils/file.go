package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// DataFolder : Pasta raíz para dados
const DataFolder = "data/"

// InimigoFolder : pasta de dados dos inimigos
const InimigoFolder = "inimigos/"

// PersonagensFolder : pasta de dados dos inimigos
const PersonagensFolder = "personagens/"

// ItensFolder : pasta de dados dos itens
const ItensFolder = "itens/"

// JSONToVar : Carrega um arquivo JSON e coloca numa variável. Ele usa o ponteiro da variável e o caminho do arquivo JSON.
func JSONToVar(v interface{}, path string) {
	jsonFile, _ := os.Open(path)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &v)
}

// CreatePath : Cria um caminho passando o tipo de objeto que quer pegar e o id do objeto
func CreatePath(objeto int8, id int8) string {
	var path string = DataFolder
	if objeto == 1 {
		path += InimigoFolder
	}
	if id == 1 {
		path += "rato.json"
	}
	return path
}
