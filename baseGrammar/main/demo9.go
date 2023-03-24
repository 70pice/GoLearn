package main

import (
	"baseGrammar/main/funk"
	"fmt"
)

type Person interface {
	getName() string
}

type ClassMate struct {
	age  int
	name string
}

func (classMate *ClassMate) getName() string {
	return classMate.name
}

func main() {
	var p Person = &ClassMate{18, "ice"}
	fmt.Println(p.getName())
	ice := funk.Ice
	fmt.Println(ice)
}
