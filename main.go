package main

import (
	"github.com/angelo848/esquenta2-golang/http"
	"github.com/angelo848/esquenta2-golang/model"
	uuid "github.com/satori/go.uuid"
)

func main() {
	produto1 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Carrinho",
	}

	produto2 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Boneca",
	}

	prodcts := model.Products{}
	prodcts.Add(produto1)
	prodcts.Add(produto2)

	server := http.NewWebServer()
	server.Products = &prodcts
	server.Serve()
}
