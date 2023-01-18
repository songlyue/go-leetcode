package main

func main() {
	ints := []int{3, 2, 2, 3}
	print(removeElement(ints, 3))
}

/**
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。
例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。

*/

func removeElement(nums []int, val int) int {
	fastIndex, slowIndex := 0, 0
	for ; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}
