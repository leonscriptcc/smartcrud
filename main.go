package main

import (
	"log"
	"smartcrud/model"
	"smartcrud/smartcrud"
)

func main() {
	g := smartcrud.InitGen("/Users/leonscript/GolandProjects/smartcrud/test", "model", model.User{})
	err := g.GenerateCRUD()
	if err != nil {
		log.Println(err)
	}
}
