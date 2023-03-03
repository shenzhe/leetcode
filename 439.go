package code

import "math"

var tmp439 []int
var count439 int

func Execuate_439(nums []int) int {
	return reversePairs(nums)
}

func reversePairs(nums []int) int {
	l := len(nums)
	tmp439 = make([]int, l)
	count439 = 0
	MergeSort439(nums)
	return count439
}

func MergeSort439(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	// fmt.Println(nums)
	sort439(nums, 0, l-1)
}

func sort439(nums []int, lo, hi int) {
	//单个元素不用排序
	if lo == hi {
		return
	}
	//防止整型溢出
	mid := lo + (hi-lo)/2
	// fmt.Println("mid", lo, hi, mid)
	//排左边的
	sort439(nums, lo, mid)
	// fmt.Println("left", nums[lo:mid])
	//排右边的
	sort439(nums, mid+1, hi)
	//合并左右两边有序为一个有序数组
	merge439(nums, lo, mid, hi)
}

func merge439(nums []int, lo, mid, hi int) {
	//把数据存入临时数组，以便交换数据
	for i := lo; i <= hi; i++ {
		tmp439[i] = nums[i]
	}
	// fmt.Println("tmp", tmp)

	//双指针大小比较，进行交换
	s, m := lo, mid+1
	// fmt.Println(s, m, lo, hi, mid)
	end := mid + 1
	for k := lo; k <= mid; k++ {
		for end <= hi && nums[end] <= math.MaxInt/2 && nums[k] > nums[end]*2 {
			end++
		}
		count439 += end - mid - 1
	}
	//遍历数组
	for i := lo; i <= hi; i++ {
		// fmt.Println("merge", s, m, lo, hi, mid, i)
		if s == mid+1 {
			//左边的数组已合并完, 如果右边还有剩余，全部merger到nums
			// fmt.Println("change1", i, m, tmp[m])
			nums[i] = tmp439[m]
			m++
		} else if m == hi+1 {
			//右边的数组已合并完, 如果左边还有剩余，全部merger到nums
			// fmt.Println("change2", i, s, tmp[s])
			nums[i] = tmp439[s]
			s++
		} else if tmp439[s] > tmp439[m] {
			//如果左边的大，把右边小的存入nums对应的位置，右边向往走一格
			// fmt.Println("change3", i, m, tmp[m])
			nums[i] = tmp439[m]
			m++
		} else {
			//如果左边的小，把左边的存入nums对应的位置，左边向前走
			// fmt.Println("change4", i, s, tmp[s])
			nums[i] = tmp439[s]
			s++
		}
	}

	// fmt.Println("nums", nums)
}
