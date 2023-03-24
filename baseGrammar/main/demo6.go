package main

import "fmt"

func demo6() {
	arrA := []int{1, 1, 3, 4, 5}

	for i := 1; i < len(arrA); i++ {
		fmt.Println(arrA[i])
	}

	if arrA[1] == 2 {
		fmt.Print("在执行if")
	} else if arrA[1] == 1 {
		fmt.Println("在执行else if")
	} else {
		fmt.Print("ok")
	}

	switch arrA[1] {
	case 1:
		fmt.Println("123")
		fallthrough
	case 2:
		fmt.Println("321")
	case 3:
		fmt.Println("345")
	default:
		fmt.Println("465")
	}

}
