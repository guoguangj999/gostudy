package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 4, 9, 7, 0, 8, 5, 6}
	fmt.Print(twoSum(nums))

	var nums2 = []int{2, 7, 11, 10}

	target := 9

	fmt.Print(twoSum2(nums2, target))
}

func twoSum(nums []int) []int {

	for i := 0; i < len(nums); i++ {

		for j := i + 0; j < len(nums); j++ {

			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp

			}
		}
	}

	return nums
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		n := target - nums[i]

		if val, ok := m[n]; ok {
			return []int{val, i}
		}
		m[nums[i]] = i
	}
	return nil
}
