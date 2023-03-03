package code

func Execuate_283(nums []int) {
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	l := len(nums)
	j := 0
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	for j < l {
		nums[j] = 0
		j++
	}
}
