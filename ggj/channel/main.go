package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//prod()
	//wor()
	//fmt.Printf("sum is %d", sum)
	go test()
	//我们在调用test函数是在前面添加关键字 go
	//表明我们开启一个协程来运行这个函数
	//协程不会影响我们主线程，所以会同时运行main中的程序
	for i := 0; i < 5; i++ {

		fmt.Println("main()  hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

var cha = make(chan int)
var sum = 0

func prod() {

	for i := 0; i <= 100; i++ {
		cha <- i
	}
	close(cha)
	//time.Sleep(time.Millisecond)

}

func wor() {

	fmt.Print("协程开始执行")
	go func() {
		for {
			num := <-cha
			sum = sum + num
		}

	}()
	fmt.Print("-----------------")
	for i := 0; i <= 100; i++ {
		cha <- i
	}

}

func test() {
	for i := 0; i < 5; i++ {

		fmt.Println("test()  hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
