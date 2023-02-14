package main

import (
	"fmt"
)

func main() {
	x := 121
	fmt.Print(huiwen(x))
}

func huiwen(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	//反转后的数
	reverNumber := 0

	for x > reverNumber {
		//取余获取后半部分的数字
		reverNumber = reverNumber*10 + x%10
		//前半部分数
		x /= 10
	}
	return x == reverNumber/10 || x == reverNumber
}
