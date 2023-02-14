package main

import (
	"fmt"
	"runtime"
)

func Sum() {
	var xx []int
	for i := 0; i < 40000000; i++ {
		if i%2 == 0 {
			xx = append(xx, i)
		}
	}
}

func main() {

	Sum()

	num := runtime.NumCPU() //这里得到当前所有的cpu数量
	fmt.Printf("couNum=%v\n", num)

	runtime.GOMAXPROCS(num - 1) //这里我们预留一个cpu给其他程序用
	fmt.Println("OK")

}
