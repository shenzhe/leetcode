package code

import (
	"fmt"
	"reflect"
	"testing"
)

func Test105(t *testing.T) {

}

func Test106(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := Execuate106(inorder, postorder)

	PostTraverseTreeNode(root)
	InTraverseTreeNode(root)
}

func TestO33(t *testing.T) {
	trueOder := []int{1, 2, 3, 4, 5}
	falseOrder := []int{1, 6, 3, 2, 5}

	if !Execuate_o33(trueOder) {
		t.Errorf("order is true")
	}

	if Execuate_o33(falseOrder) {
		t.Errorf("order is false")
	}
}

func Test315(t *testing.T) {
	nums := []int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41}
	result := []int{10, 27, 10, 35, 12, 22, 28, 8, 19, 2, 12, 2, 9, 6, 12, 5, 17, 9, 19, 12, 14, 6, 12, 5, 12, 3, 0, 10, 0, 7, 8, 4, 0, 0, 4, 3, 2, 0, 1, 0}

	ans := Execuate_315(nums)

	if !reflect.DeepEqual(ans, result) {
		t.Errorf("test 315 error")
	}
}

func Test439(t *testing.T) {
	nums := []int{12, 2, 1, 17, 19, 10, 5, 23, 7, 20, 10, 17, 22, 15, 9, 18, 12, 12, 16, 16, 17, 8, 11, 19, 2, 21, 5, 19, 22, 9, 17, 24, 8, 8, 16, 5, 2, 25, 1, 0, 3, 24, 25, 0, 11, 7, 19, 0, 5, 16, 17, 4, 19, 20, 20, 0, 14, 4, 16, 15, 11, 15, 20, 11, 17, 13, 3, 18, 12, 6, 10, 25, 12, 6, 18, 6, 19, 19, 18, 13, 21, 9, 17, 1, 1, 2, 10, 15, 24, 24, 22, 7, 10, 23, 15, 9, 1, 23, 22, 15, 3, 16, 23, 25, 8, 18, 0, 5, 1, 12, 9, 0, 25, 0, 13, 11, 22, 5, 3, 13, 10, 17, 14, 24, 23, 1, 8, 1, 21, 18, 2, 16, 21, 21, 5, 3, 19, 8, 23, 6, 6, 3, 2, 4, 13, 2, 4, 14, 9, 17, 23, 18, 4, 23, 5, 13, 25, 10, 9, 14, 3, 9, 11, 5, 14, 18, 0, 10, 13, 5, 19, 17, 24, 25, 4, 8, 16, 14, 3, 24, 18, 2, 17, 22, 4, 11, 18, 9, 9, 7, 10, 4, 24, 0, 7, 0, 6, 15, 18, 13, 14, 20, 22, 17, 22, 15, 17, 9, 10, 17, 13, 0, 22, 22, 23, 2, 21, 18, 6, 10, 10, 15, 14, 4, 4, 18, 21, 15, 0, 18, 14, 0, 2, 24, 6, 10, 1, 8, 25, 20, 13, 20, 13, 20, 5, 21, 21, 9, 19, 8, 9, 9, 5, 17, 18, 18, 20, 5, 17, 18, 3, 7, 21, 6, 0, 8, 3, 3, 1, 11, 0, 21, 6, 15, 11, 10, 13, 6, 7, 21, 7, 1, 1, 14, 15, 20, 2, 8, 21, 25, 19, 12, 18, 16, 0, 4, 10, 19, 14, 23, 6, 17, 2, 15, 19, 4, 13, 8, 14, 4, 15, 21, 4, 23, 20, 3, 18, 0, 12, 14, 14, 19, 0, 21, 18, 21, 17, 13, 9, 20, 17, 25, 17, 21, 16, 22, 4, 1, 13, 20, 15, 9, 7, 18, 18, 7, 22, 8, 18, 1, 13, 0, 24, 8, 12, 16, 1, 3, 6, 23, 16, 24, 5, 0, 1, 25, 3, 16, 9, 4, 24, 1, 11, 24, 9, 16, 11, 0, 2, 20, 16, 0, 1, 6, 19, 22, 12, 3, 23, 21, 4, 20, 1, 0, 18, 24, 10, 0, 12, 21, 17, 23, 0, 13, 1, 25, 9, 19, 0, 13, 21, 23, 6, 24, 25, 16, 9, 8, 16, 2, 22, 23, 3, 7, 16, 25, 11, 18, 19, 4, 11, 1, 25, 22, 9, 11, 14, 9, 3, 16, 8, 5, 11, 12, 15, 15, 19, 15, 15, 7, 17, 24, 18, 9, 8, 20, 23, 18, 17, 7, 8, 19, 23, 9, 13, 4, 17, 23, 21, 19, 11, 22, 22, 9, 3, 19, 23, 11, 2, 23, 8, 8, 21, 15, 1, 25, 7, 6, 14, 6, 7, 11, 3, 2, 11, 14, 10, 24, 3, 8, 10, 1, 18, 4, 6, 16, 12, 18, 12, 6, 5, 25, 24, 25, 7, 12, 17, 19, 15, 8, 23, 7, 6, 11, 6, 16, 14, 15, 13, 18, 5, 9, 21, 24, 8, 17, 25, 21, 22, 19, 24, 9, 9, 25, 21, 6, 25, 24, 3, 15, 20, 19, 13, 7, 13, 3, 0, 11, 2, 3, 23, 4, 14, 13, 7, 14, 3, 2, 18, 6, 1, 24, 19, 11, 6, 22, 9, 20, 3, 15, 23, 14, 18, 11, 11, 0, 2, 14, 21, 1, 12, 8, 8, 22, 10, 25, 20, 15, 22, 15, 21, 4, 19, 23, 5, 20, 4, 10, 17, 9, 7, 8, 11, 7, 10, 2, 18, 5, 24, 4, 16, 22, 13, 0, 11, 6, 19, 8, 21, 23, 24, 14, 19, 6, 3, 1, 17, 25, 22, 9, 14, 12, 15, 2, 24, 23, 17, 3, 3, 3, 6, 11, 20, 11, 0, 12, 17, 0, 3, 12, 24, 5, 13, 11, 19, 5, 2, 5, 12, 20, 19, 23, 2, 14, 23, 19, 4, 6, 15, 12, 2, 24, 17, 18, 9, 18, 4, 12, 20, 17, 19, 21, 16, 15, 13, 0, 17, 10, 23, 22, 10, 8, 20, 6, 4, 13, 11, 0, 3, 1, 5, 19, 17, 23, 17, 10, 10, 7, 4, 1, 20, 21, 23, 21, 21, 25, 2, 1, 8, 22, 4, 10, 16, 9, 15, 12, 12, 7, 3, 10, 14, 11, 9, 0, 7, 1, 1, 18, 23, 16, 6, 4, 20, 17, 18, 20, 17, 22, 8, 19, 6, 8, 14, 23, 14, 14, 15, 3, 24, 19, 16, 18, 14, 3, 6, 10, 8, 22, 12, 6, 8, 5, 3, 20, 10, 15, 19, 17, 8, 10, 7, 22, 0, 5, 19, 18, 16, 22, 24, 6, 18, 19, 19, 21, 1, 22, 14, 0, 24, 1, 20, 21, 7, 2, 11, 13, 10, 9, 7, 13, 15, 22, 2, 17, 4, 1, 4, 22, 22, 7, 18, 3, 12, 12, 7, 6, 20, 15, 25, 8, 13, 7, 5, 1, 25, 12, 1, 25, 16, 3, 23, 25, 9, 22, 4, 11, 16, 21, 20, 15, 17, 16, 13, 14, 20, 5, 23, 9, 0, 6, 3, 21, 2, 7, 2, 22, 7, 5, 8, 17, 14, 17, 8, 18, 21, 22, 14, 8, 15, 2, 10, 24, 0, 10, 23, 11, 16, 22, 5, 5, 19, 20, 14, 2, 19, 3, 25, 5, 10, 14, 22, 3, 5, 10, 20, 22, 16, 17, 22, 15, 23, 10, 0, 21, 17, 20, 3, 15, 0, 13, 17, 2, 10, 20, 8, 24, 5, 6, 19, 9, 4, 25, 11, 19, 10, 3, 24, 0, 10, 10, 9, 21, 16, 25, 6, 20, 11, 7, 17, 20, 10, 9, 22, 19, 21, 7, 0, 4, 11, 1, 9, 18, 18, 3, 1, 25, 5, 1, 20, 13, 2, 7, 19, 10, 13, 25, 3, 23, 13, 5, 10, 15, 11, 15, 22, 9, 10, 8, 18, 0}
	result := 124430
	ans := Execuate_439(nums)

	if result != ans {
		t.Errorf("test439 error, want=%d, but=%d", result, ans)
	}
}

/*
[-2,5,-1]
-2
2
*/

func Test327(t *testing.T) {
	nums := []int{-2, 5, -1}
	lower := -2
	upper := 2
	ans := Execuate_327(nums, lower, upper)
	result := 3

	if result != ans {
		t.Errorf("test327 error, want=%d, but=%d", result, ans)
	}
}

func Test230(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7} // -1 表示 null
	root := BuildBST(nums)
	//前
	fmt.Println("pre====")
	PreTraverseTreeNode(root)
	//中
	fmt.Println("in====")
	InTraverseTreeNode(root)
	//后
	fmt.Println("post====")
	PostTraverseTreeNode(root)
	k := 3
	result := nums[k-1]
	ans := Execuate_230(root, k)

	if result != ans {
		t.Errorf("test327 error, want=%d, but=%d", result, ans)
	}
}

func Test215(t *testing.T) {
	// nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6} // -1 表示 null
	// result := 4
	// k := 4
	type cases struct {
		nums   []int
		k      int
		result int
	}
	nums := []cases{
		{
			nums:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:      4,
			result: 4,
		},
		{
			nums:   []int{3, 2, 1, 5, 6, 4},
			k:      2,
			result: 5,
		},
	}
	for _, item := range nums {
		// fmt.Println(item.nums, item.k, item.result)
		ans := Execuate_215(item.nums, item.k)

		if item.result != ans {
			t.Errorf("test327 error, want=%d, but=%d", result, ans)
		}
	}
}

func Test380(t *testing.T) {
	rs := NewRandomizedSet()
	//type[1=insert, 2= remove, 3=Random], val, result
	cases := [][3]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 2, 1},
		{2, 2, 1},
		{2, 2, 0},
		{3, 1, 1},
		{1, 3, 1},
		{1, 4, 1},
		{2, 1, 1},
		{1, 5, 1},
		{1, 6, 1},
		{3, 5, 1},
	}
	for _, v := range cases {
		switch v[0] {
		//insert
		case 1:
			ans := rs.Insert(v[1])
			ret := 0
			if ans {
				ret = 1
			}
			if ret != v[2] {
				t.Errorf("insert error, wait=%d, but=%d", v[2], ret)
			}
		case 2:
			ans := rs.Remove(v[1])
			ret := 0
			if ans {
				ret = 1
			}
			if ret != v[2] {
				t.Errorf("insert error, wait=%d, but=%d", v[2], ret)
			}
		case 3:
			for i := 0; i < v[1]; i++ {
				ret, idx := rs.GetRandomAndIdx()
				// rs.PrintAllMapping()
				ans := rs.Get(idx)
				if ret != ans {
					t.Errorf("random error, wait=%d, but=%d(idx=%d)", ret, ans, idx)
				}
			}
		}
	}
}

func Test167(t *testing.T) {
	type item struct {
		numbers []int
		target  int
		ans     []int
	}
	cases := []item{
		{
			[]int{2, 7, 11, 15}, 9, []int{1, 2},
		},
		{
			[]int{2, 3, 4}, 6, []int{1, 3},
		},
		{
			[]int{-1, 0}, -1, []int{1, 2},
		},
	}

	for _, v := range cases {
		ans := Execuate_167(v.numbers, v.target)

		if !reflect.DeepEqual(ans, v.ans) {
			t.Errorf("numbers=%v, target=%d, want=%v, but=%v", v.numbers, v.target, v.ans, ans)
		}
	}
}

// "babbad"
func Test5(t *testing.T) {
	cases := map[string]string{
		"babbad": "abba",
		"babad":  "bab",
		"cbbd":   "bb",
		"":       "",
		"d":      "d",
	}
	for k, v := range cases {
		ans := Execuate_5(k)
		fmt.Printf("s=%s, want=%v, ans=%v\n", k, v, ans)
		if v != ans {
			t.Errorf("s=%s, want=%v, ans=%v", k, v, ans)
		}
	}
}

func Test344(t *testing.T) {
	cases := map[string]string{
		"abcd": "dcba",
		"ab":   "ba",
		"b":    "b",
		"":     "",
	}
	for k, v := range cases {
		ans := Execuate_344(k)
		fmt.Printf("s=%s, want=%v, ans=%v\n", k, v, ans)
		if v != ans {
			t.Errorf("s=%s, want=%v, ans=%v", k, v, ans)
		}
	}
}

func Test283(t *testing.T) {
	type item struct {
		nums []int
		ans  []int
	}
	cases := []item{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{0},
			[]int{0},
		},
	}

	for _, item := range cases {
		Execuate_283(item.nums)
		if !reflect.DeepEqual(item.nums, item.ans) {
			t.Errorf("want=%v, but=%v", item.ans, item.nums)
		}
	}
}

func TestInterveiw_0(t *testing.T) {
	nums := []int{0, 3, 0, 4, 3, 4, 3}
	ans := Execuate_interview_0(nums)
	fmt.Printf("ans=%v", ans)
}

func Test151(t *testing.T) {
	cases := map[string]string{
		"abcd":                   "abcd",
		"  abcd":                 "abcd",
		"abcd  ":                 "abcd",
		"  abcd  ":               "abcd",
		"abcd efgh":              "efgh abcd",
		"abcd  efgh":             "efgh abcd",
		" abcd  efgh":            "efgh abcd",
		" abcd  efgh ":           "efgh abcd",
		"abcd  efgh   ":          "efgh abcd",
		"a good   a  example":    "example a good a",
		"   a good   a  example": "example a good a",
	}
	for k, v := range cases {
		ans := Execuate_151(k)
		fmt.Printf("s=%s, want=%s, but=%s\n", k, v, ans)
		if ans != v {
			t.Errorf("s=%s, want=%s, but=%s", k, v, ans)
		}
	}
}

func Test797(t *testing.T) {
	graph := [][]int{
		{3, 1},
		{4, 6, 7, 2, 5},
		{4, 6, 3},
		{6, 4},
		{7, 6, 5},
		{6},
		{7},
		{},
	}
	Execuate_797(graph)
}

func Test1443(t *testing.T) {
	// n := 7
	// edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	// // hasApple := []bool{false, false, true, false, false, true, false}
	// hasApple := []bool{false, false, true, false, true, true, false}
	n := 4
	edges := [][]int{{0, 1}, {1, 2}, {0, 3}}
	hasApple := []bool{true, true, true, true}
	fmt.Printf("n=%d\n", n)
	ans := Execuate_1443(n, edges, hasApple)
	fmt.Printf("ans=%d\n", ans)
}

func Test2049(t *testing.T) {
	parents := []int{-1, 3, 3, 5, 7, 6, 0, 0}
	// parents := []int{-1, 2, 0, 2, 0}
	Execuate_2049(parents)
}

func Test79(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCEE"

	ans := Execuate_79(board, word)
	fmt.Printf("ans=%t\n", ans)
}

func Test210(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	ans := Execuate_210(numCourses, prerequisites)
	fmt.Printf("ans=%v\n", ans)

}

func Test785(t *testing.T) {
	type tcase struct {
		graph [][]int
		ans   bool
	}
	cases := []tcase{
		// {
		// 	[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
		// 	false,
		// },
		// {
		// 	[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
		// 	true,
		// },
		{
			[][]int{{3, 4, 6}, {3, 6}, {3, 6}, {0, 1, 2, 5}, {0, 7, 8}, {3}, {0, 1, 2, 7}, {4, 6}, {4}, {}},
			true,
		},
	}

	for _, item := range cases {
		ans := Execuate_785(item.graph)
		if ans != item.ans {
			t.Errorf("graph=%v, want=%t, but=%t", item.graph, item.ans, ans)
		}
	}
}

func Test886(t *testing.T) {
	type tcase struct {
		n        int
		dislikes [][]int
		ans      bool
		is_bsf   bool
	}
	cases := []tcase{
		{
			4,
			[][]int{{1, 2}, {1, 3}, {2, 4}},
			true,
			true,
		},
	}

	for _, item := range cases {
		ans := Execuate_886(item.n, item.dislikes, item.is_bsf)
		if ans != item.ans {
			t.Errorf("dislike=%v, want=%t, but=%t", item.dislikes, item.ans, ans)
		}
	}
}

func Test130(t *testing.T) {
	type tcase struct {
		board [][]byte
		ans   [][]byte
	}
	cases := []tcase{
		{
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
			// [][]byte{{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'}, {'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'}, {'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'}, {'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X'}, {'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'}, {'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'O'}, {'X', 'O', 'X', 'O', 'O', 'X', 'X', 'O', 'X', 'O'}, {'X', 'X', 'O', 'X', 'X', 'O', 'X', 'O', 'O', 'X'}, {'O', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O'}, {'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'}},
			// [][]byte{{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'}, {'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'}, {'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'}, {'O', 'O', 'O', 'O', 'O', 'O', 'X', 'O', 'O', 'X'}, {'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'}, {'X', 'O', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'O'}, {'X', 'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'O'}, {'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'X'}, {'O', 'O', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O'}, {'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'}},
		},
	}
	for _, item := range cases {
		fmt.Printf("start=%v\n", item.board)
		Execuate_130_3(item.board)
		fmt.Printf("end=%v\n", item.board)
		fmt.Printf("ans=%v\n", item.ans)
	}
}

func Test990(t *testing.T) {
	type tcase struct {
		equations []string
		ans       bool
	}
	cases := []tcase{
		{
			[]string{"a==b", "b!=a"},
			false,
		},
		{
			[]string{"b==a", "a==b"},
			true,
		},
		{
			[]string{"a==b", "b!=c", "c==a"},
			false,
		},
	}
	for _, item := range cases {
		ans := Execuate_990(item.equations)
		if ans != item.ans {
			t.Errorf("equations=%v, want=%t, but=%t", item.equations, item.ans, ans)
		}
	}
}

func Test1361(t *testing.T) {
	type tcase struct {
		n          int
		leftChild  []int
		rightChild []int
		ans        bool
	}
	cases := []tcase{
		{
			3,
			[]int{3, -1, 1, -1},
			[]int{-1, -1, 0, -1},
			false,
		},
		{
			4,
			[]int{1, -1, 3, -1},
			[]int{2, -1, -1, -1},
			true,
		},
		{
			4,
			[]int{1, -1, 3, -1},
			[]int{2, 3, -1, -1},
			false,
		},
		{
			2,
			[]int{1, 0},
			[]int{-1, -1},
			false,
		},
		{
			6,
			[]int{1, -1, -1, 4, -1, -1},
			[]int{2, -1, -1, 5, -1, -1},
			false,
		},
	}
	for _, item := range cases {
		ans := Execuate_1361(item.n, item.leftChild, item.rightChild)
		if ans != item.ans {
			t.Errorf("item=%v, want=%t, but=%t", item, item.ans, ans)
		}
	}
}

func Test1584(t *testing.T) {
	type tcase struct {
		points [][]int
		ans    int
	}
	cases := []tcase{
		{
			[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			20,
		},
		{
			[][]int{{3, 12}, {-2, 5}, {-4, 1}},
			18,
		},
		{
			[][]int{{-8988, -8416}, {-5299, 7210}, {-5132, -6801}, {-501, 4378}, {-6429, -2747}, {-6707, 7699}, {4136, -6086}, {6311, -5534}, {-4114, -9305}, {-8189, -9624}, {9644, -4646}, {-8520, 3180}, {9018, -2556}, {-5445, 4930}, {1549, 7421}, {-455, -7400}, {3578, 1030}, {-9676, -7740}, {6823, -2933}, {8599, 1252}, {-9122, 6121}, {-4331, -5318}, {8399, 4967}, {1454, -7459}, {-8826, -6339}, {-1601, -6797}, {-9473, 247}, {1092, -4900}, {5946, -2760}, {-4227, 9007}, {-2034, -1698}, {1647, 7561}, {1353, -9845}, {8132, 9102}, {8012, 6644}, {8496, -4241}, {-4084, 1708}, {2145, -5993}, {4171, 7629}, {3700, 8578}, {-4532, 9034}, {1504, -8983}, {-2943, -2501}, {9505, -8352}, {2729, 3954}, {-1834, -1837}, {-1083, -7265}, {5475, -1753}, {2684, 589}, {1269, -4238}, {-2160, 9037}, {-5638, 6135}, {5638, 2211}, {5235, 8446}, {-3559, -3221}, {372, 9599}, {1091, 3439}, {-1939, 2607}, {7559, 7687}, {2411, -9285}, {774, 6395}, {8933, -4090}, {1729, -7803}, {7707, 221}, {192, 6442}, {2666, -8556}, {-9394, -7278}, {-3436, 6198}, {-6875, -5577}, {-2867, -8426}, {-8793, -3110}, {-7341, -169}, {5555, -1718}, {-9725, -4156}, {-2756, 7213}, {1413, -6421}, {3261, 2341}, {-4749, 9258}, {-8220, 1700}, {-8242, 6550}, {6618, -1170}, {6851, 8241}, {5619, 3580}, {-8944, 1442}, {-990, 9490}, {9526, 2764}, {2867, 9679}, {-1257, -8846}, {-2790, 7544}, {-8419, -4198}, {-3594, 4537}, {7874, -1396}, {3646, -7597}, {7350, 2628}, {9478, -2709}, {-4623, -5005}, {9044, -3746}, {-6955, -5785}, {-2175, 5650}, {4367, -3521}, {8504, -1616}, {-4546, -2365}, {8841, 5989}, {-5191, 1178}, {9692, 125}, {-593, 1298}, {8350, -2744}, {-603, 9710}, {8983, -3423}, {-6393, -1619}, {5439, -2874}, {4686, -9336}, {8776, 4842}, {-1654, -2478}, {-1325, 107}, {-2511, -7791}, {-8477, -5472}, {-7026, -427}, {-2426, -5032}, {2, 7761}, {3937, -4022}, {5355, 9532}, {337, -2619}, {-2166, -4865}, {-314, -2231}, {-6403, -356}, {8370, 3473}, {-4575, 1726}, {-3380, -7740}, {-8538, 3338}, {-9502, -1065}, {-6139, 8975}, {-5613, -3139}, {-739, -4537}, {-5365, 329}, {7868, 3052}, {-6290, 7374}, {-8647, 5963}, {9889, -2064}, {9974, -4514}, {5313, -7159}, {-7519, 1076}, {7748, -3701}, {-8180, 3306}, {9686, 2669}, {9247, -5331}, {7282, -7061}, {-3077, 9371}, {-3418, 3846}, {5637, 4391}, {-4230, -121}, {5453, 9099}, {1959, 6109}, {5093, -8844}, {5054, 8312}, {3659, -940}, {8363, 9969}, {-101, -264}, {-11, 5271}, {9460, 5886}, {494, -4020}, {-8877, 8164}, {-2303, -5103}, {-7535, -6647}, {-5321, 8877}, {-5981, 6099}, {5451, 2469}, {-1469, -7877}, {-5688, -9714}, {8648, -7318}, {-4535, 6785}, {2974, 5314}, {-6017, -5672}, {6821, 1176}, {5458, 3798}, {-2684, 6453}, {813, 8999}, {-6524, 8952}, {-4414, 6571}, {3109, -9498}, {-3787, 5517}, {1510, -7545}, {5376, 8396}, {4069, 7141}, {-562, 397}, {7604, -6051}, {7237, -9629}, {-1970, 5925}, {-9338, -8166}, {-6285, 7404}, {7655, 8779}, {-5772, -101}, {6310, 6307}, {-4812, 2604}, {7992, 7534}, {1100, -1035}, {1318, -8732}, {8995, -9691}, {-8415, -7895}, {-5136, -6000}, {-3625, -330}, {1030, -4247}, {2787, 1385}, {-624, -7305}, {-1976, -2812}, {-741, -9974}, {-6193, -4721}, {1199, 157}, {8295, 9380}, {-5658, -5301}, {-6675, -4990}, {-2181, -701}, {5192, 408}, {-1692, 6007}, {-3271, -8525}, {6862, -9175}, {-2055, 1786}, {5905, -486}, {3032, 1347}, {1581, 4776}, {-4249, -949}, {2943, 3646}, {-8158, 1645}, {-7362, 2848}, {3488, -3259}, {3483, -7734}, {-1535, -8670}, {-5878, -7635}, {-5536, 3947}, {-2244, 658}, {5875, 4050}, {7989, -1189}, {8642, 7381}, {-6640, -8838}, {4971, -2124}, {-6951, -8377}, {-452, -8449}, {1789, -3908}, {-8879, -3281}, {8668, -1666}, {2412, 3345}, {-4813, 5466}, {6476, 361}, {6271, 4086}, {-1623, -3600}, {-5234, -7773}, {3992, 7081}, {5964, 7402}, {-5158, -307}, {-9346, 5117}, {-2788, 1219}, {-9691, -7467}, {3983, -3384}, {9258, -5179}, {7457, -3511}, {8945, -603}, {-3818, -9181}, {7617, -6696}, {2759, 3932}, {1528, -8928}, {7434, -3165}, {-1069, -4960}, {3207, -5549}, {-6276, -1351}, {3744, -9299}, {-6972, 3162}, {-6886, -2824}, {-9271, 7476}, {-1595, 1728}, {-29, -1217}, {7249, 8212}, {1146, -6591}, {1784, 6986}, {-1585, 1461}, {-6850, 1592}, {2703, -8838}, {-7391, 2755}, {-5065, -5565}, {-4756, 7472}, {-3846, 4104}, {-5041, -1126}, {-4896, -9019}, {-3624, 6080}, {4077, -1332}, {8627, -5382}, {-5786, 9566}, {-4567, 4490}, {2253, -6215}, {-579, -9047}, {-647, -1579}, {-2278, -7088}, {6855, -3423}, {6192, 6193}, {-3962, -3285}, {-785, 6981}, {6796, 2665}, {6749, -2647}, {7149, 840}, {7234, 2152}, {559, -2138}, {6132, -4945}, {-8386, 1981}, {-4861, 413}, {2423, -1170}, {7790, -831}, {-7736, 3443}, {-2634, 4789}, {860, 365}, {3460, -6434}, {-1297, -2600}, {-6190, -8501}, {-7956, 9131}, {-3060, -4059}, {-9644, 7930}, {-2873, 6790}, {5874, -6554}, {-9417, -1207}, {-456, 4639}, {-4566, 9926}, {-9438, 2825}, {-6042, -3183}, {-6495, -3877}, {7744, -6847}, {-7017, 7854}, {1848, 6927}, {-4389, 1563}, {-4795, 6535}, {-5241, 4654}, {6237, -3216}, {7194, -2613}, {-4326, -5007}, {-9698, -778}, {6241, 5525}, {6461, 6649}, {-8533, 5329}, {1165, 6283}, {-4292, -7563}, {7594, 3596}, {-3525, -750}, {2025, -4290}, {2056, 3713}, {-1580, -2}, {-3026, -4448}, {913, 1398}, {1784, 9855}, {6823, 5176}, {-4101, -7117}, {9378, -820}, {5188, -7520}, {6441, -5630}, {9220, -2304}, {-700, -5775}, {-1778, -8855}, {-5798, -4711}, {1363, -5175}, {9020, 4459}, {326, 7169}, {6069, 2083}, {-1523, 1319}, {-5172, -1917}, {9427, 8576}, {-8713, -5436}, {6950, 2613}, {7137, -7896}, {5799, -4602}, {6194, 1863}, {1324, 3743}, {-9168, -6443}, {1465, -1789}, {-65, -5349}, {-6439, -1075}, {-7931, 7102}, {3888, -4675}, {1411, -9841}, {-6485, -9322}, {9441, -8060}, {4728, 6773}, {-9411, 1317}, {6466, 3161}, {1892, -5184}, {-8238, -4791}, {2143, 5460}, {4833, 1765}, {-98, -8624}, {-6431, 1285}, {7949, -6954}, {-2162, 1485}, {-6691, -4410}, {6405, 6575}, {7663, -6223}, {-9712, 3091}, {977, 4378}, {-4846, 6846}, {1136, -2407}, {-9413, 8690}, {8173, -8524}, {-2463, -8684}, {5856, -527}, {-9484, 5483}, {5568, 2532}, {-3325, 640}, {-6906, 8347}, {8943, 1320}, {1972, 9204}, {9670, 6174}, {-2539, 8811}, {8071, -2298}, {4322, 9261}, {-824, -640}, {667, -7559}, {2935, 3765}, {961, 1933}, {-4225, 554}, {-2672, -3510}, {5510, -4440}, {4814, -6211}, {5334, 115}, {4862, 3075}, {-1938, -7453}, {5563, 8333}, {4903, 4369}, {-1962, -6210}, {-4275, -2512}, {-6366, -3346}, {9252, 5797}, {7291, 1919}, {-8155, -9318}, {9531, -3092}, {-8120, 7513}, {-8421, 306}, {3297, -3803}, {9593, -9650}, {8578, 4628}, {5431, 3966}, {-9600, 675}, {4979, -5283}, {1616, 215}, {-252, -733}, {1846, 102}, {-4586, 3178}, {-7218, -2268}, {-252, 4357}, {6752, 9575}, {4045, 7027}, {2872, 145}, {-2515, 8818}, {7343, 7036}, {-9071, -8979}, {1567, -8532}, {5710, 4639}, {-7649, -4757}, {-2316, -3250}, {-740, -8523}, {-8934, -6745}, {9941, 2407}, {1952, -7321}, {4152, -448}, {-9594, 1085}, {-2192, -7473}, {9257, 4488}, {8369, -1238}, {-3452, -8993}, {-6954, 5505}, {-312, 4659}, {-7728, 6193}, {8092, -6396}, {-909, 5334}, {4077, 9062}, {-2916, -2665}, {-9831, 7040}, {-9131, 2894}, {-1036, -8186}, {-4364, -5071}, {-5484, 1140}, {-9514, 6979}, {-6770, 5050}, {-9931, 6683}, {-2080, 5690}, {-4430, -5829}, {8014, -7985}, {5821, -8714}, {5457, 9374}, {-7217, 5224}},
			361721,
		},
	}
	for _, item := range cases {
		ans := Execuate_1584(item.points)
		if ans != item.ans {
			t.Errorf("item=%v, want=%d, but=%d", item, item.ans, ans)
		}
	}

	// for _, item := range cases {
	// 	ans := Execuate_1584_1(item.points)
	// 	if ans != item.ans {
	// 		t.Errorf("item=%v, want=%d, but=%d", item, item.ans, ans)
	// 	}
	// }
}

func Test743(t *testing.T) {
	type tcase struct {
		times [][]int
		n     int
		k     int
		ans   int
	}

	cases := []tcase{
		{
			[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			4,
			2,
			2,
		},
		{
			[][]int{{1, 2, 1}},
			2,
			1,
			1,
		},
		{
			[][]int{{1, 2, 1}},
			2,
			2,
			-1,
		},
	}
	for _, item := range cases {
		ans := Execuate_743(item.times, item.n, item.k)
		if ans != item.ans {
			t.Errorf("times=%v, want=%d, but=%d\v", item, item.ans, ans)
		}
	}
}

func Test1631(t *testing.T) {
	type tcase struct {
		heights [][]int
		ans     int
	}

	cases := []tcase{
		{
			[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			2,
		},
		{
			[][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}},
			1,
		},
	}
	for _, item := range cases {
		ans := Execuate_1631(item.heights)
		if ans != item.ans {
			t.Errorf("times=%v, want=%d, but=%d\v", item, item.ans, ans)
		}
	}
}

func Test1514(t *testing.T) {
	//n = 3, edges = {{0,1},{1,2},{0,2}}, succProb = {0.5,0.5,0.2}, start = 0, end = 2
	type tcase struct {
		n        int
		edges    [][]int
		succProb []float64
		start    int
		end      int
		ans      float64
	}

	cases := []tcase{
		{
			3,
			[][]int{{0, 1}, {1, 2}, {0, 2}},
			[]float64{0.5, 0.5, 0.3},
			0,
			2,
			0.3,
		},
		{
			3,
			[][]int{{0, 1}, {1, 2}, {0, 2}},
			[]float64{0.5, 0.5, 0.2},
			0,
			2,
			0.25,
		},
		{
			3,
			[][]int{{0, 1}},
			[]float64{0.5},
			0,
			2,
			0,
		},
	}
	for _, item := range cases {
		ans := Execuate_1514(item.n, item.edges, item.succProb, item.start, item.end)
		if ans != item.ans {
			t.Errorf("times=%v, want=%f, but=%f\v", item, item.ans, ans)
		}
	}
}

func Test46(t *testing.T) {
	type tcase struct {
		nums []int
		ans  [][]int
	}

	cases := []tcase{
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			[]int{5, 4, 6, 2},
			[][]int{{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}},
		},
	}
	for _, item := range cases {
		ans := Execuate_46(item.nums)
		if !reflect.DeepEqual(ans, item.ans) {
			t.Errorf("times=%v, want=%v, but=%v\v", item, item.ans, ans)
		}
	}
}

func Test51(t *testing.T) {
	type tcase struct {
		n   int
		ans [][]string
	}

	cases := []tcase{
		{
			4,
			[][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
	}
	for _, item := range cases {
		ans := Execuate_51(item.n)
		if !reflect.DeepEqual(ans, item.ans) {
			t.Errorf("times=%v, want=%v, but=%v\v", item, item.ans, ans)
		}
	}
}

func Test140(t *testing.T) {
	type tcase struct {
		s        string
		wordDict []string
		ans      []string
	}

	cases := []tcase{
		{
			"catsanddog",
			[]string{"cat", "cats", "and", "sand", "dog"},
			[]string{"cats and dog", "cat sand dog"},
		},
	}
	for _, item := range cases {
		ans := Execuate_140(item.s, item.wordDict)
		if !reflect.DeepEqual(ans, item.ans) {
			t.Errorf("times=%v, want=%v, but=%v\v", item, item.ans, ans)
		}
	}
}

func Test17(t *testing.T) {
	type tcase struct {
		digits string
		ans    []string
	}

	cases := []tcase{
		{
			"23",
			[]string{"cat", "cats", "and", "sand", "dog"},
		},
	}
	for _, item := range cases {
		ans := Execuate_17(item.digits)
		if !reflect.DeepEqual(ans, item.ans) {
			t.Errorf("times=%v, want=%v, but=%v\v", item, item.ans, ans)
		}
	}
}