package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	slice := make([]int, 2)

outer:
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				slice[0] = i
				slice[1] = j
				break outer
			}
		}
	}

	return slice
}

func main() {
	nums := []int{3, 3}
	fmt.Println(twoSum(nums, 6))
}
