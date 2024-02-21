package models

type Produto struct {
	ID         int
	Nome       string  `validate:"required"`
	Descricao  string  `validate:"required"`
	Preco      float64 `validate:"required"`
	Quantidade int     `validate:"required"`
}
