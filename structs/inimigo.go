package structs

// Inimigo : Struct que vai receber os dados do inimigo
type Inimigo struct {
	ID          int8   `json:"id"`
	Nome        string `json:"nome"`
	Vida        int16  `json:"vida"`
	Ataque      int8   `json:"ataque"`
	Defesa      int8   `json:"defesa"`
	Experiencia int16  `json:"experiencia"`
	Item        int8   `json:"item"`
}
