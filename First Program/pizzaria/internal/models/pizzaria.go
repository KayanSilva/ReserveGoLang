package models

type Pizza struct {
	ID      int      `json:"id"`
	NOME    string   `json:"nome"`
	PRECO   float32  `json:"preco"`
	Reviews []Review `json:"reviews,omitempty"`
}
