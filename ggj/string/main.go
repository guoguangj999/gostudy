package main

import "fmt"

func main() {

	test("IVXLM")
	right := 8
	left := 0
	mid := (right-left)>>1 + left

	fmt.Println("-------", mid)
}

func test(s string) {
	var replacemap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := range s {
		//fmt.Println(i)
		fmt.Println(replacemap[s[i]])
	}

}
