package code

func Execuate_o83(nums []int) [][]int {
	return permute_o83(nums)
}

func permute_o83(nums []int) [][]int {
	l := len(nums)
	ans_o83 := make([][]int, 0)
	visited_083 := make([]bool, l)
	path := make([]int, 0)
	traverse_o83(nums, path, l, &ans_o83, &visited_083)
	return ans_o83
}

func traverse_o83(nums []int, path []int, end int, ans_o83 *[][]int, visited_083 *[]bool) {

	if len(path) == end {
		//加入答案列表
		newpath := make([]int, len(path))
		copy(newpath, path)
		*ans_o83 = append(*ans_o83, newpath)
		return
	}

	for i := 0; i < end; i++ {
		//已选择
		visited := *visited_083
		if visited[i] {
			continue
		}
		path = append(path, nums[i])
		visited[i] = true
		traverse_o83(nums, path, end, ans_o83, visited_083)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
