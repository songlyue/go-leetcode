package main

import (
	"fmt"
	"math"
)

/**
209. 长度最小的子数组

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

*/

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	arrayLen := minSubArrayLen(7, nums)
	fmt.Println(arrayLen)
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	start := 0
	end := 0
	sum := 0
	ans := math.MaxInt
	if len(nums) == 0 {
		return 0
	}
	for end < n {
		sum += nums[end]
		for sum >= target {
			length := min(ans, end-start+1)
			ans = length
			sum -= nums[start]
			start++
		}
		end++
	}

	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
