package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários"))
}

func raiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página raiz"))
}

func main() {
	/*
		HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB
		CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)
		O cliente faz uma requisição onde o servidor vai processar essa requisição e vai devolver uma resposta
		Request - Response

		Rotas - Identificar o tipo de mensagem que está sendo enviada e
		identificar que tipo de processamento o servidor vai fazer em cima da mensagem
		Rotas
		URI - Identificador do recurso
		Método - GET, POST, PUT, DELETE
	*/

	http.HandleFunc("/", raiz)

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
