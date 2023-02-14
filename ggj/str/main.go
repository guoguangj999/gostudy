package main

import (
	"fmt"
	"reflect"
)

func main() {

	strs := "qwer"
	index := 2
	fmt.Println(strs[:index])

	v1 := "(){}[]"

	for i := range v1 {
		fmt.Println(v1[i], reflect.TypeOf(v1[i]))
	}

	var stack = []byte{}

	stack = append(stack, '(')

	top := stack[:len(stack)-1]
	fmt.Println(top)

}
