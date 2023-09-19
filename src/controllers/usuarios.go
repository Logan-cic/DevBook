package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter,  r *http.Response){
	w.Write([]byte("Criando Usuário"))
}

func BuscarUsuarios(w http.ResponseWriter,  r *http.Response){
	w.Write([]byte("Buscando todos os usuários"))
}

// func BusUsuario(w http.ResponseWriter,  r *http.Response){
// 	w.Write([]byte("Buscando todos os usuários"))
// }

// func CriarUsuario(w http.ResponseWriter,  r *http.Response){
// 	w.Write([]byte("Criando Usuário"))
// }

// func CriarUsuario(w http.ResponseWriter,  r *http.Response){
// 	w.Write([]byte("Criando Usuário"))
// }