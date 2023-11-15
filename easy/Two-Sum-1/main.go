//    Given an array of integers nums and an integer target,
//    return indices of the two numbers such that they add up to target.
//    You may assume that each input would have exactly one solution,
//    and you may not use the same element twice.
//    You can return the answer in any order.

//    Example 1:
//    Input: nums = [2,7,11,15], target = 9
//    Output: [0,1]
//    Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

//    Example 2:
//    Input: nums = [3,2,4], target = 6
//    Output: [1,2]

//    Example 3:
//    Input: nums = [3,3], target = 6
//    Output: [0,1]

//    Constraints:
//    2 <= nums.length <= 104
//    -109 <= nums[i] <= 109
//    -109 <= target <= 109
//    Only one valid answer exists.

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

func findTarget(nums []int, target int, hashMap map[int][]int) []int {
	for i, num := range nums {
		remaining := target - num
		values, isMatch := hashMap[remaining]

		// no match
		if !isMatch {
			continue
		}
		// match
		if !(values[0] == i) {
			return []int{i, values[0]}
		}
		// match, but is itself, therefore gets the next occurrence if it existed
		if values[0] == i && len(values) > 1 {
			return []int{i, values[1]}
		}
	}

	return []int{}
}

func toMap(nums []int) map[int][]int {
	dataMap := make(map[int][]int)
	for i := 1; i < len(nums); i++ {
		dataMap[nums[i]] = append(dataMap[nums[i]], i)
	}

	return dataMap
}
