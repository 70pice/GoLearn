package main

import "fmt"

func demo5() {
	// 声明切片的方式
	var arrA = [5]int{11, 22, 33, 44, 55}
	var sliceA = arrA[1:3]
	fmt.Print(len(sliceA))
	fmt.Print(cap(sliceA))

	mapA := make(map[string]int)

	mapB := map[string]int{
		"srar": 1,
		"qwe":  2,
	}
	fmt.Print(mapA, mapB)
	str := "Golang"
	var p *string = &str
	fmt.Println(p)

}
