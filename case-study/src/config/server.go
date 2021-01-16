package config

import (
	"fmt"
	"net/http"
)

func Server() {
	configHandler()

	fmt.Print("Servidor está rodando na porta 1337")

	_ = http.ListenAndServe(":1337", nil) //DefaultServerMux
}
