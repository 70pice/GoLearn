package main

import "fmt"

// 虽然编译器能通过，但是定义数组一定要规定长度！不规定长度的都是异类
var arr [][]int

// 这样的定义才是正确的
var arr2 = [][]int{{1, 2}, {3, 4}}

var arr3 = [5]int{1, 2, 3}

func f() {
	fmt.Println(arr3)
	fmt.Println(arr2)
	for i := 0; i < len(arr3); i++ {
		fmt.Print(arr3[i])
	}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[0]); j++ {
			fmt.Println(arr2[i][j])
		}
	}

	fmt.Println("执行函数前的arr3为")
	fmt.Println(arr3)
	arr3 = function(arr3)
	fmt.Println("执行函数后的arr3为")
	fmt.Println(arr3)
}

func function(arr [5]int) [5]int {
	for i := 0; i < len(arr); i++ {
		arr[i] += 2
	}
	// 在go中数组是值传递，而不是传入了一个地址有点意思哈
	return arr
}
