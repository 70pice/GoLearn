package main

import "fmt"

func demo4() {
	// 数组的定义方式
	var arrA [5]int
	// 如果想声明数组并且初始化
	var arrB = [5]int{1, 2, 3, 4, 5}
	// 或者使用 := 赋值运算符
	arrC := [5]int{1, 2, 3, 4, 5}
	for i := 1; i < len(arrA); i++ {
		arrA[i] = i
	}
	fmt.Println(arrB)
	fmt.Println(arrC)
	// 注意，在go中，数组是一种值，在参数传递的过程中，会拷贝一份相同的数组，而不会传递地址
	// 注意，在go中，数组的 == 比较的是值而不是地址
}
