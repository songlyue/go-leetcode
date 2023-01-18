package main

func main() {
	print(isValid("["))
}

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

输入：s = "()[]{}"
输出：true
*/

func isValid(s string) bool {
	stack := make([]byte, 0)
	hash := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, hash[s[i]])
		} else if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
