package main

/*
*
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
*/
func main() {

}
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	set := make(map[int]bool, 0)
	for _, i := range nums1 {
		set[i] = true
	}
	for _, i := range nums2 {
		if set[i] {
			set[i] = false
			res = append(res, i)
		}
	}
	return res
}
