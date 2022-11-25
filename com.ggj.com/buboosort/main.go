package main

import "fmt"

func main() {

	nums := []int{3, 7, 4, 8, 1, 9}

	fmt.Println(buboo(nums))

}

func buboo(nums []int) []int {
	n := len(nums)

	//最后一次比较的位置
	lastExchangeIndex := 0

	//每次比较都到这里为止
	sortBoder := n - 1
	for i := 0; i < n-1; i++ {
		isSort := true
		for j := 0; j < sortBoder; j++ {
			if nums[j] > nums[j+1] {
				isSort = false //发生数据位置变换，变为false
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp

				lastExchangeIndex = j

			}
		}

		sortBoder = lastExchangeIndex
		if isSort {
			// s := 0
			// s++
			// fmt.Println(s)
			break
		}
	}

	return nums
}
