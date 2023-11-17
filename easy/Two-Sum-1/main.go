package main

import "fmt"

func main() {
	resultA := twoSum([]int{2, 7, 11, 15}, 9)
	resultB := twoSum([]int{3, 2, 4}, 6)
	resultC := twoSum([]int{3, 3}, 6)
	fmt.Println("resultA: ", resultA)
	fmt.Println("resultB: ", resultB)
	fmt.Println("resultC: ", resultC)
}

func twoSum(nums []int, target int) []int {
	// edge case 1: only two elements
	if len(nums) == 2 {
		return []int{0, 1}
	}

	return findTarget(nums, target, toMap(nums))
}

func findTarget(nums []int, target int, dataMap map[int]int) []int {
	currentIndex := -1
	targetIndex := -1
	for i := 0; i < len(nums)-1; i++ {
		remaining := target - nums[i]
		value, isFound := dataMap[remaining]

		if !isFound {
			continue
		} else {
			// is self
			if value == i {
				continue
			}

			currentIndex = i
			targetIndex = value
			break
		}
	}

	return []int{currentIndex, targetIndex}
}

func toMap(nums []int) map[int]int {
	dataMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dataMap[nums[i]] = i
	}

	return dataMap
}
