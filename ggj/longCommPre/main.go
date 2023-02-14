package main

import "fmt"

//最长公共前缀
func main() {

	var strs = []string{"flower", "flow", "flight"}

	fmt.Println(longPrefix(strs))

}

func longComm(s1 string, s2 string) string {
	length := min(len(s1), len(s2))
	index := 0

	for index < length && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

func longPrefix(s []string) string {
	prefix := s[0]
	for i := 1; i < len(s); i++ {
		prefix = longComm(prefix, s[i])

		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
