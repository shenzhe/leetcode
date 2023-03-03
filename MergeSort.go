package code

import "fmt"

var count_ms int

func MergeSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	res := sort1(nums)
	fmt.Println("count", count_ms)
	return res

}

func sort1(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	mid := l / 2
	leftNums := sort1(nums[:mid])
	rightNums := sort1(nums[mid:])
	// fmt.Println(mid, leftNums, rightNums)
	return merge(leftNums, rightNums)
}

func merge(leftNums []int, rightNums []int) []int {
	ll := len(leftNums)
	rl := len(rightNums)

	res := make([]int, 0, ll+rl)
	i, j := 0, 0
	//当left或right有一个遍历完成，则跳出for
	for i < ll && j < rl {
		if leftNums[i] < rightNums[j] {
			res = append(res, leftNums[i])
			i++
		} else {
			res = append(res, rightNums[j])
			j++
		}
		count_ms++
	}
	//merge可能剩余的leftnums或rightnums
	if i < ll {
		count_ms += ll - i
		res = append(res, leftNums[i:]...)
	}
	if j < rl {
		count_ms += rl - 1
		res = append(res, rightNums[j:]...)
	}
	// fmt.Println(res)
	return res
}

var tmp []int

func MergeSort2(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	count_ms = 0
	tmp = make([]int, l)
	// fmt.Println(nums)
	sort2(nums, 0, l-1)
	fmt.Println("count", count_ms)
}

func sort2(nums []int, lo, hi int) {
	//单个元素不用排序
	if lo == hi {
		return
	}
	//防止整型溢出
	mid := lo + (hi-lo)/2
	// fmt.Println("mid", lo, hi, mid)
	//排左边的
	sort2(nums, lo, mid)
	// fmt.Println("left", nums[lo:mid])
	//排右边的
	sort2(nums, mid+1, hi)

	//合并左右两边有序为一个有序数组
	merge2(nums, lo, mid, hi)
}

func merge2(nums []int, lo, mid, hi int) {
	//把数据存入临时数组，以便交换数据
	for i := lo; i <= hi; i++ {
		count_ms++
		tmp[i] = nums[i]
	}
	// fmt.Println("tmp", tmp)

	//双指针大小比较，进行交换
	s, m := lo, mid+1
	// fmt.Println(s, m, lo, hi, mid)

	//遍历数组
	for i := lo; i <= hi; i++ {
		// fmt.Println("merge", s, m, lo, hi, mid, i)
		if s == mid+1 {
			//左边的数组已合并完, 如果右边还有剩余，全部merger到nums
			// fmt.Println("change1", i, m, tmp[m])
			nums[i] = tmp[m]
			m++
		} else if m == hi+1 {
			//右边的数组已合并完, 如果左边还有剩余，全部merger到nums
			// fmt.Println("change2", i, s, tmp[s])
			nums[i] = tmp[s]
			s++
		} else if tmp[s] > tmp[m] {
			//如果左边的大，把右边小的存入nums对应的位置，右边向往走一格
			// fmt.Println("change3", i, m, tmp[m])
			nums[i] = tmp[m]
			m++
		} else {
			//如果左边的小，把左边的存入nums对应的位置，左边向前走
			// fmt.Println("change4", i, s, tmp[s])
			nums[i] = tmp[s]
			s++
		}
		count_ms++
	}

	// fmt.Println("nums", nums)
}
