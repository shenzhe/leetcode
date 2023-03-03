package code

import "math"

/*
n个数组里，里面的数字都 小于n, 求超过1个的元素的个数，并返回
[1,2,3,4,3,4,3]
[2,1,3,4,3,4,6]
[3,1,2,4,3,4,6]
[4,1,2,3,3,4,6]
[3,1,2,3,4,4,6]



{
	3:2,
	4:2
}
*/

func Execuate_interview_0(nums []int) map[int]int {
	res := make(map[int]int)
	for _, num := range nums {
		num = int(math.Abs(float64(num)))
		if nums[num] <= 0 {
			if _, ok := res[num]; ok {
				res[num]++
			} else {
				res[num] = 2
			}
		} else {
			nums[num] *= -1
		}
	}
	if _, ok := res[0]; ok {
		if nums[0] == 0 {
			res[0]--
		}
	}
	return res
}

/**

 */
