package main

import "fmt"

// 定义a为int类型的数据，go是有类型的变成语言
var a int

// 定义变量的时候就可以赋予初始值
var b int = 1

// 可以自动识别变量的类型
var c = 1

// 还有一种简便的写法
// 如果直接这么写会报错，这是因为 b := 123  == 》 var b int  + b = 123 这两条代码，而 b = 123 这样的代码不允许在函数体外执行

func test() {

	var e byte = 1
	var f float32 = 1.1
	var g float64 = 3.4
	var msg string = "321"
	helloWorld()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	d := 123
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(msg)
}
