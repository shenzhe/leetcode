package code

var tmp327 []int64
var count327 int

func Execuate_327(nums []int, lower, upper int) int {
	return countRangeSum(nums, lower, upper)
}

func countRangeSum(nums []int, lower, upper int) int {
	l := len(nums)

	if l < 1 {
		return 0
	}

	if l == 1 {
		if lower == upper && lower == nums[0] {
			return 1
		}
		return 0
	}

	count327 = 0
	//构建一个前缀和数组
	newNums := make([]int64, l+1)
	//构建一个临时数组
	tmp327 = make([]int64, l+1)
	for i := 0; i < l; i++ {
		newNums[i+1] = int64(nums[i]) + newNums[i]
	}
	// fmt.Println(nums, newNums)
	sort327(newNums, 0, l, lower, upper)
	return count327
}

func sort327(nums []int64, lo, hi, lower, upper int) {
	//单个元素不用排序
	if lo == hi {
		return
	}
	//防止整型溢出
	mid := lo + (hi-lo)/2
	// fmt.Println("mid", lo, hi, mid)
	//排左边的
	sort327(nums, lo, mid, lower, upper)
	// fmt.Println("left", nums[lo:mid])
	//排右边的
	sort327(nums, mid+1, hi, lower, upper)
	//合并左右两边有序为一个有序数组
	merge327(nums, lo, mid, hi, lower, upper)
}

func merge327(nums []int64, lo, mid, hi, lower, upper int) {
	//把数据存入临时数组，以便交换数据
	for i := lo; i <= hi; i++ {
		tmp327[i] = nums[i]
	}
	// fmt.Println("tmp", tmp)

	//双指针大小比较，进行交换
	s, m := lo, mid+1
	// fmt.Println(s, m, lo, hi, mid)
	start, end := mid+1, mid+1
	for k := lo; k <= mid; k++ {
		//小于lower的值
		for start <= hi && nums[start]-nums[k] < int64(lower) {
			start++
		}
		for end <= hi && nums[end]-nums[k] <= int64(upper) {
			end++
		}
		count327 += end - start
	}
	//遍历数组
	for i := lo; i <= hi; i++ {
		// fmt.Println("merge", s, m, lo, hi, mid, i)
		if s == mid+1 {
			//左边的数组已合并完, 如果右边还有剩余，全部merger到nums
			// fmt.Println("change1", i, m, tmp[m])
			nums[i] = tmp327[m]
			m++
		} else if m == hi+1 {
			//右边的数组已合并完, 如果左边还有剩余，全部merger到nums
			// fmt.Println("change2", i, s, tmp[s])
			nums[i] = tmp327[s]
			s++
		} else if tmp327[s] > tmp327[m] {
			//如果左边的大，把右边小的存入nums对应的位置，右边向往走一格
			// fmt.Println("change3", i, m, tmp[m])
			nums[i] = tmp327[m]
			m++
		} else {
			//如果左边的小，把左边的存入nums对应的位置，左边向前走
			// fmt.Println("change4", i, s, tmp[s])
			nums[i] = tmp327[s]
			s++
		}
	}

	// fmt.Println("nums", nums)
}
