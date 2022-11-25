package main

import "fmt"

//全局变量声明了，可以不使用
var (
	name string
)

//main函数，程序的入口
func main() {
	fmt.Println("富贵是个傻逼！")
	name = "杨卓"
	fmt.Printf("name:%s", name) //%s 是占位符
	//简短变量声明，只可以在函数中使用
	s1 := "shabi"
	fmt.Print(s1)
}
