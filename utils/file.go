package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// JSONToVar : get json
func JSONToVar(v interface{}) {
	jsonFile, _ := os.Open("data/inimigos/rato.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &v)
}
