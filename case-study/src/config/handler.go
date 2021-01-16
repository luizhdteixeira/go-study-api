package config

import (
	"case-study/src/api"
	"fmt"
	"net/http"
)

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Bem-vindo!")

}

func configHandler() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/books", api.FindAllBooks)

}
