package main

import "fmt"

/*
*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
*/
func main() {
	nums := []int{-4, -1, 0, 3, 10}
	squares := sortedSquares(nums)
	fmt.Println(squares)
}

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	k := len(nums) - 1

	res := make([]int, len(nums))
	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[k] = nums[left] * nums[left]
			left++
		} else {
			res[k] = nums[right] * nums[right]
			right--
		}
		k--
	}
	return res
}
