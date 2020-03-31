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

// JSONToVar : Carrega um arquivo JSON e coloca numa variável. Ele usa o ponteiro da variável e o caminho do arquivo JSON.
func JSONToVar(v interface{}, path string) {
	jsonFile, _ := os.Open(path)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &v)
}
