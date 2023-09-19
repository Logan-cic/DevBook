package rotas

import "net/http"

//ROTA representa todas as rotas da API
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}