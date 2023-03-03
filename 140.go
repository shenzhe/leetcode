package code

func Execuate_140(s string, wordDict []string) []string {
	return wordBreak(s, wordDict)
}

var ans_140 []string
var path_140 string

func wordBreak(s string, wordDict []string) []string {
	l := len(s)
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	ans_140 = make([]string, 0)
	path_140 = ""
	traverse_140(s, wordMap, 0, l)
	return ans_140
}

func traverse_140(s string, wordMap map[string]bool, idx, l int) {
	if idx == l {
		ans_140 = append(ans_140, path_140)
		return
	}

	for i := idx + 1; i <= l; i++ {
		//命中dict
		if _, ok := wordMap[s[idx:i]]; ok {
			if len(path_140) > 0 {
				path_140 += " "
			}
			path_140 += s[idx:i]
			traverse_140(s, wordMap, i, l)
			nl := len(path_140) - 1 - i + idx
			if nl < 0 {
				nl = 0
			}
			path_140 = path_140[:nl]
		}
	}
}
