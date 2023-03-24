package main

// 在Go中一个文件夹通常就为一个包，其中特殊的是package main，它声明了main.go所在的包

import "fmt"

// fmt是标准库，用来处理输入与输出

// 通常采用func来定义一个函数，main为主函数
func helloWorld() {
	fmt.Println("hello world")
}

// 此时如果执行go run demo1.go
// 等价于执行go build demo1.go + ./demo1

// 在go的程序中，只有main包可以包含main函数
// 在go程序中，一个main包中有且只能有一个main函数
