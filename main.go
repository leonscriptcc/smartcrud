package main

import (
	"log"
	"reflect"
)

func main() {
	var test interface{}
	test = 1
	t := reflect.TypeOf(test)
	log.Println(t.Kind().String())
}
