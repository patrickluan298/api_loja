package repositories

import (
	"database/sql"
	"fmt"

	"github.com/api_loja/models"
)

func AdicionaProduto(nome string, descricao string, preco float64, quantidade int) error {
	Connection()
	defer db.Close()

	_, err = db.Exec("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)", nome, descricao, preco, quantidade)
	if err != nil {
		return fmt.Errorf("erro ao adicionar produto: %v", err)
	}

	return nil
}

func BuscaTodosProdutos() ([]models.Produto, error) {
	Connection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		return nil, fmt.Errorf("erro ao listar produtos: %v", err)
	}

	produtos := []models.Produto{}

	for rows.Next() {
		var produto models.Produto

		err := rows.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear produtos: %v", err)
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}

func EditaProduto(id string) (models.Produto, error) {
	Connection()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM produtos WHERE id = $1", id)

	produto := models.Produto{}

	err := row.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Produto{}, fmt.Errorf("produto com ID %s não encontrado", id)
		}
		return models.Produto{}, err
	}

	return produto, nil
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) error {
	Connection()
	defer db.Close()

	_, err = db.Exec("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5", nome, descricao, preco, quantidade, id)
	if err != nil {
		return fmt.Errorf("falha ao atualizar produto: %v", err)
	}

	return nil
}

func DeletaProduto(id string) error {
	Connection()
	defer db.Close()

	_, err = db.Exec("DELETE FROM produtos WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("falha ao excluir produto produto: %v", err)
	}

	return nil
}
