package main

import (
	"errors"
	"fmt"
)

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("fileName shouldn`t be nil")
	}
	return nil
}
func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("happen error")
			ret = -1
		}
	}()
	arr := []int{1, 2, 3}
	return arr[index]
}
func main() {
	//_, err := os.Open("fileName.txt")
	//if err != nil {
	//	fmt.Print(err)
	//}
	defer get(1)
	get(2)
	get(3)
}
