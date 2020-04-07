package structs

// Personagem : Struct que vai receber os dados do personagem
type Personagem struct {
	ID     int8   `json:"id"`
	Nome   string `json:"nome"`
	Vida   int16  `json:"vida"`
	Ataque int8   `json:"ataque"`
	Defesa int8   `json:"defesa"`
}
