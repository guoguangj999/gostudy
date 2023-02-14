package main

import "fmt"

//type关键字用于定义类型
type Student struct {
	Name  string
	Score int
}

//结构体方法，方法中可以使用结构体变量
func (s Student) Study() {
	s.Score += 10
}

//结构体指针方法，方法中可以使用结构体指针变量
func (s *Student) Study1() {
	s.Score += 10
}

func main() {

	stu := Student{
		Name:  "张三",
		Score: 60,
	}

	stu1 := &stu
	fmt.Println(stu.Score) //60

	stu.Study()
	fmt.Println(stu.Score)

	stu.Study1()
	fmt.Println(stu.Score)

	stu1.Study()
	fmt.Println(stu1.Score)

	stu1.Study1()
	fmt.Println(stu1.Score)

}
