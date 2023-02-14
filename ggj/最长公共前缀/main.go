package main

import "fmt"

func main() {
	var strs = []string{"flower", "flow", "flight"}
	//fmt.Print(commonPre(strs))
	fmt.Print(ZhongHeng(strs))
}

//返回公共的部分
func longCom(s1, s2 string) string {
	length := min2(len(s1), len(s2))
	index := 0
	for index < length && (s1[index] == s2[index]) {
		index++
	}
	return s1[:index]
}

//对字符串数组中的字符串进行一一比较找出公共的部分
func commonPre(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = longCom(prefix, strs[i])

		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

//找出最小字符串长度
func min2(x, y int) int {

	if x > y {
		return y
	}
	return x

}

//纵向比较
func ZhongHeng(strs []string) string {

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {

			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}

		}

	}
	return strs[0]
}
