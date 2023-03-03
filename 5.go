package code

func Execuate_5(s string) string {
	return longestPalindrome(s)
}

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	//babad
	start, end, maxLen := 0, 0, 0
	res := ""
	for i := 0; i < l; i++ {
		//奇数
		start, end = i, i
		for start >= 0 && end < l && s[start] == s[end] {
			start--
			end++
		}
		if end-start > maxLen {
			maxLen = end - start
			res = s[start+1 : end]
		}
		// fmt.Println(i, "odd", start, end, res, maxLen)
		//偶数
		start, end = i, i+1
		for start >= 0 && end < l && s[start] == s[end] {
			start--
			end++
		}
		if end-start > maxLen {
			maxLen = end - start
			res = s[start+1 : end]
		}
		// fmt.Println(i, "even", start, end, res, maxLen)
	}

	return res

}
