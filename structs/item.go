package structs

// Item : Struct que vai receber os dados do item
type Item struct {
	ID        int8   `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}

// ItemPocao : id da pocao
const ItemPocao int8 = 0
