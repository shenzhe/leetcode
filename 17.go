package code

import "fmt"

func Execuate_17(digits string) []string {
	return letterCombinations(digits)
}

func letterCombinations(digits string) []string {
	mapping := map[string][]byte{
		"2": {'a', 'b', 'c'},
		"3": {'d', 'e', 'f'},
		"4": {'g', 'h', 'i'},
		"5": {'j', 'k', 'l'},
		"6": {'m', 'n', 'o'},
		"7": {'p', 'q', 'r', 's'},
		"8": {'t', 'u', 'v'},
		"9": {'w', 'x', 'y', 'z'},
	}
	fmt.Printf("mapping=%v", mapping)
	ans := make([]string, 0)
	path := make([]byte, 0)
	traverse_17(digits, 0, path, mapping, &ans)
	return ans
}

func traverse_17(digits string, idx int, path []byte, mapping map[string][]byte, ans *[]string) {
	if idx >= len(digits) {
		return
	}
	for _, v := range mapping[string(digits[idx])] {
		path = append(path, v)
		// fmt.Printf("path1=%s, v=%s,idx=%d\n", string(path), string(v), idx)
		traverse_17(digits, idx+1, path, mapping, ans)
		if len(path) == len(digits) {
			*ans = append(*ans, string(path))
			// path = path[:0]
		}
		path = path[:len(path)-1]
		// fmt.Printf("path2=%s, v=%s, idx=%d\n", string(path), string(v), idx)

	}

}
