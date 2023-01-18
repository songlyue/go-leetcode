package main

/*
*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，
该函数将返回左旋转两位得到的结果"cdefgab"。
输入: s = "abcdefg", k = 2
输出: "cdefgab"
*/
func main() {
	sb := "abcdefg"
	print(reverseLeftWords(sb, 2))
}
func reverseLeftWords(s string, n int) string {
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	b := []byte(s)
	//abcdefg
	reverse(&b, 0, n-1)      // bacdefg
	reverse(&b, n, len(b)-1) //bagfedc
	reverse(&b, 0, len(b)-1) // cdefgab
	return string(b)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}

}
