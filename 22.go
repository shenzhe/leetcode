package code

func Execuate_22(n int) []string {
	return generateParenthesis(n)
}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	path := ""
	ans_22 := make([]string, 0)
	traverse_22(n, n, path, &ans_22)
	return ans_22
}

func traverse_22(left, right int, path string, ans_22 *[]string) {
	if right < left {
		path = ""
		return
	}

	if left < 0 || right < 0 {
		path = ""
		return
	}

	if left == 0 && right == 0 {
		// fmt.Printf("path=%s, ans=%v\v", path, ans_22)
		*ans_22 = append(*ans_22, path)
		path = ""
		return
	}

	//放一个(
	path += "("
	traverse_22(left-1, right, path, ans_22)
	path = path[:len(path)-1]

	path += ")"
	traverse_22(left, right-1, path, ans_22)
	path = path[:len(path)-1]
}
