package main

import "fmt"

func main() {

	slice := []int{1, 2, 3}

	slice[1] = 100

	slice = append(slice, 4, 5, 6)
	fmt.Println(len(slice), cap(slice), slice)

	//遍历切片
	for idx, v := range slice {
		slice[1] = 99
		fmt.Println(idx, v)
	}

	//切片截取
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := slice1[2:5]

	slice2[0] = 100

	fmt.Println(len(slice1), cap(slice1), slice1)
	fmt.Println(len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22)
	slice2[0] = 200
	fmt.Println(len(slice1), cap(slice1), slice1)
	fmt.Println(len(slice2), cap(slice2), slice2)

	slice3 := make([]int, 2)
	slice3[0] = 1
	slice3[1] = 2
	testslice(slice3)
	fmt.Println(len(slice3), cap(slice3), slice3)

}

func testslice(slice []int) {
	slice[0] = 100
	slice[1] = 200
	fmt.Println(len(slice), cap(slice), slice)

}
