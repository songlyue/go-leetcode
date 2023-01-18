package main

import "fmt"

/**
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

输入：s = "abcdefg", k = 2
输出："bacdfeg"
*/

func main() {
	ints := []int{1, 2, 3, 4}
	i := ints[1:3]
	fmt.Printf("%v", i)
}
func reverseStr(s string, k int) string {
	res := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(res) {
			reverse(res[i : i+k])
		} else {
			reverse(res[i:])
		}
	}
	return string(res)
}

func reverse(b []byte) {
	l := 0
	r := len(b) - 1
	for l < r {
		b[l], b[r] = b[r], b[l]
		r--
		l++
	}
}
