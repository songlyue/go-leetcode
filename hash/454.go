package main

/*
*
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
*/
func main() {
	print(1 - 3)
}
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	var res int
	for _, i := range nums1 {
		for _, i2 := range nums2 {
			m[i+i2]++
		}
	}
	for _, i := range nums3 {
		for _, i2 := range nums4 {
			// m[-(i+i2)]存在才加
			res += m[-(i + i2)]
		}
	}
	return res
}
