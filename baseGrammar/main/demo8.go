package main

import "fmt"

type Student struct {
	age  int
	name string
}

func (stu *Student) hello(persion string) string {
	return fmt.Sprintf("hello %s,i am %s", persion, stu.name)
}

func demo8() {
	stu := &Student{11, "ice"}
	fmt.Println(stu.hello("alice"))

}
