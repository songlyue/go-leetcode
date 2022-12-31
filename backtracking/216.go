package main

import "fmt"

/*
*
找出所有相加之和为n 的k个数的组合，且满足下列条件：

只使用数字1到9
每个数字最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combination-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	combinationSum3(9, 45)
	fmt.Println(res1)
}

var (
	path1 []int
	res1  [][]int
)

func combinationSum3(k int, n int) [][]int {
	path1 = make([]int, 0, k)
	res1 = make([][]int, 0)
	dfs3(n, k, 0, 1)
	return res1
}

func dfs3(targetSum, k, sum, startIndex int) {
	// 剪枝
	if sum > targetSum {
		return
	}

	if len(path1) == k {
		if sum == targetSum {
			temp := make([]int, k)
			copy(temp, path1)
			res1 = append(res1, temp)
		}
		return
	}
	for i := startIndex; i <= 9; i++ {
		sum += i
		path1 = append(path1, i)
		dfs3(targetSum, k, sum, i+1)
		sum -= i
		path1 = path1[:len(path1)-1]
	}
}
