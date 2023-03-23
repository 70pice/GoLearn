package main

import (
	"demo/calc"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	sum := calc.Add(1, 2)
	sub := calc.Sub(1, 2)
	fmt.Println("hello world")
	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(r)
}
