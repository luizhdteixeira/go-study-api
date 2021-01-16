package api

import (
	"case-study/src/mock"
	"encoding/json"
	"net/http"
)

func FindAllBooks(writer http.ResponseWriter, request *http.Request) {
	_ = json.NewEncoder(writer).Encode(mock.Books)

}
