package main

import (
	"fmt"
	"sort"
)

/**
题意：给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？
找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例： 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。 满足要求的四元组集合为： [ [-1, 0, 0, 1], [-2, -1, 1, 2], [-2, 0, 0, 2] ]

#思路
*/

func main() {
	ints := []int{4, -9, -2, -2, -7, 9, 9, 5, 10, -10, 4, 5, 2, -4, -2, 4, -9, 5}
	sum := fourSum(ints, -13)
	fmt.Printf("%v", sum)
}
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for one := 0; one < len(nums)-3; one++ {
		if one > 0 && nums[one] == nums[one-1] {
			continue
		}
		for second := one + 1; second < len(nums)-2; second++ {
			if second > one+1 && nums[second] == nums[second-1] {
				continue
			}
			third := second + 1
			four := len(nums) - 1
			for third < four {
				sum := nums[one] + nums[second] + nums[third] + nums[four]
				if sum == target {
					temp := []int{nums[one], nums[second], nums[third], nums[four]}
					res = append(res, temp)
					// 继续往下走
					for third < four && nums[third] == nums[third+1] {
						third++
					}
					for third < four && nums[four] == nums[four-1] {
						four--
					}
					third++
					four--
				}
				if sum > target {
					four--
				}
				if sum < target {
					third++
				}
			}
		}
	}
	return res
}
