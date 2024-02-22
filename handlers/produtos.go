package handlers

import (
	"html/template"
	"net/http"

	"github.com/api_loja/repositories"

	"strconv"
)

// Retorna os templates .html apontados no arquivo
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos, err := repositories.BuscaTodosProdutos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = temp.ExecuteTemplate(w, "Index", todosProdutos); err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusBadRequest)
		return
	}
}

func Novo(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido, apenas POST é suportado.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao analisar o formulário", http.StatusBadRequest)
		return
	}

	camposObrigatorios := []string{"nome", "descricao", "preco", "quantidade"}
	for _, campo := range camposObrigatorios {
		if len(r.FormValue(campo)) == 0 {
			http.Error(w, "Campo '"+campo+"' é obrigatório", http.StatusBadRequest)
			return
		}
	}

	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	precoStr := r.FormValue("preco")
	quantidadeStr := r.FormValue("quantidade")

	precoConv, err := strconv.ParseFloat(precoStr, 64)
	if err != nil {
		http.Error(w, "Erro na conversão do preço", http.StatusBadRequest)
		return
	}

	quantidadeConv, err := strconv.Atoi(quantidadeStr)
	if err != nil {
		http.Error(w, "Erro na conversão da quantidade", http.StatusBadRequest)
		return
	}

	if precoConv <= 0 || quantidadeConv <= 0 {
		http.Error(w, "Preço e quantidade devem ser maiores que zero", http.StatusBadRequest)
		return
	}

	if err := repositories.AdicionaProduto(nome, descricao, precoConv, quantidadeConv); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Deletar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID não fornecido na requisição", http.StatusBadRequest)
		return
	}

	if err := repositories.DeletaProduto(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	produto, err := repositories.EditaProduto(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco := r.FormValue("preco")
	quantidade := r.FormValue("quantidade")

	if id == "" || nome == "" || preco == "" || quantidade == "" {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	idConv, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Erro na conversão do id para inteiro", http.StatusBadRequest)
		return
	}

	precoConv, err := strconv.ParseFloat(preco, 64)
	if err != nil {
		http.Error(w, "Erro na conversão do preço para float", http.StatusBadRequest)
		return
	}

	quantidadeConv, err := strconv.Atoi(quantidade)
	if err != nil {
		http.Error(w, "Erro na conversão da quantidade para inteiro", http.StatusBadRequest)
		return
	}

	if err := repositories.AtualizaProduto(idConv, nome, descricao, precoConv, quantidadeConv); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
