package main

import "fmt"

//最后一个单词的长度
func main() {
	s := "hello world "
	fmt.Println(lastWord(s))
}

func lastWord(s string) int {
	n := len(s)
	last := n - 1
	start := 0
	ans := 0
	for s[last] == 32 {
		last--
	}
	start = last

	for start >= 0 && s[start] != 32 {
		ans++
		start--
	}

	return ans

}
