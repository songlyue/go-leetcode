package main

import "sort"

/*
*
15. 三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
func main() {
	ints := []int{-1, 0, 1, 2, -1, -4}
	print(threeSum(ints))
}
func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		// 去重 如果是相同的值，那么取到的数组也是相同的，
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			// 防止内存泄露
			n2, n3 := left, right
			sum := nums[i] + nums[n2] + nums[n3]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[n2], nums[n3]})
				// 继续往下走
				if left < right && nums[left] == n2 {
					left++
				}
				if left < right && nums[right] == n3 {
					right--
				}
			}
			if sum > 0 {
				right--
			}
			if sum < 0 {
				left++
			}
		}
	}
	return res
}
