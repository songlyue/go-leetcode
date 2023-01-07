package main

/*
*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。
*/
func main() {
	print(canConstruct("aa", "ab"))
}
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[string]int)
	for _, i2 := range magazine {
		m[string(i2)]++
	}
	for _, i := range ransomNote {
		if m[string(i)] == 0 {
			return false
		}
		m[string(i)]--
	}
	return true
}
