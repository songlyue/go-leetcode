package main

/*
*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
*/
func main() {
	anagram := isAnagram("aaabbbb", "aaabbbb")
	print(anagram)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}
	for _, ch := range t {
		m[ch]--
		if m[ch] < 0 {
			return false
		}
	}
	return true
}
