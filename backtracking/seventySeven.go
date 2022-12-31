package main

import "fmt"

var (
	path []int
	res  [][]int
)

/*
*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。
*/
func main() {
	i := combine(4, 2)
	fmt.Println(i)
}

func combine(n int, k int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n, k, startIndex int) {
	if len(path) == k { // 说明已经满足了k个数的要求
		tmp := make([]int, k) // 避免闭包问题
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := startIndex; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		dfs(n, k, i+1)
		path = path[:len(path)-1]
	}
}
