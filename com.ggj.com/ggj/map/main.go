package main

import "fmt"

func main() {

	//map的声明方式

	//第一种方式
	var map1 map[string]string

	map1 = make(map[string]string, 10)

	map1["one"] = "go"
	map1["two"] = "java"
	map1["three"] = "php"

	fmt.Println(map1)

	//第二种方式
	map2 := make(map[string]string)

	map2["one"] = "fuzhou"
	fmt.Println(map2)

	//第三种方式
	score := map[string]int{
		"zhangsan": 100,
		"lisi":     99,
		"wangwu":   98,
	}
	score["who"] = 90 //往map中添加元素

	s, ok := score["who"] //map访问，s对应value的值，ok表示该key是否存在，不存在返回空

	s2, ok := score["huangkun"]

	delete(score, "lisi")
	fmt.Println(s, ok, score, s2, ok)

}
