package code

func Execuate_46(nums []int) [][]int {
	return permute(nums)
}

var ans_46 [][]int

func permute(nums []int) [][]int {
	l := len(nums)

	if l == 1 {
		return [][]int{{nums[0]}}
	}

	ans_46 = make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, l)
	backtrace_46(nums, path, used)
	return ans_46
}

func backtrace_46(nums, path []int, used []bool) {
	if len(path) == len(nums) {
		//加入ans
		newpath := make([]int, len(path))
		//需要copy新的
		copy(newpath, path)
		ans_46 = append(ans_46, newpath)
	}

	//遍历nums
	for i := 0; i < len(nums); i++ {
		if used[i] {
			//已添加到本次路径中，跳过
			continue
		}
		//加入到path
		path = append(path, nums[i])
		used[i] = true //标记为已加入
		//以nums[i]为头，进入下一组循环，达到全排列
		backtrace_46(nums, path, used)
		//以nums[i]为头的全排列已完成，标记为false，以便进入下一个数字为头的全排列
		used[i] = false
		path = path[:len(path)-1]
	}
}
