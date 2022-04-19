package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	longestStr := ""
	for i, _ := range s {
		p1 := expand(s, i, i)
		if len(p1) > len(longestStr) {
			longestStr = p1
		}

		p2 := expand(s, i, i+1)
		if len(p2) > len(longestStr) {
			longestStr = p2
		}
	}
	return longestStr
}

// 向左右扩展
func expand(s string, left int, right int) string {
	length := len(s)
	for left >= 0 && right < length && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
