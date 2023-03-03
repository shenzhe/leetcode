package code

var tmp315 [][2]int
var count315 []int

func Execuate_315(nums []int) []int {
	return countSmaller(nums)
}

func countSmaller(nums []int) []int {
	l := len(nums)
	baknums := make([][2]int, 0, l)
	for i := 0; i < l; i++ {
		baknums = append(baknums, [2]int{nums[i], i})
	}
	// fmt.Println(l, len(baknums))
	tmp315 = make([][2]int, l)
	count315 = make([]int, l)
	MergeSort315(baknums)
	// fmt.Println(count)
	return count315
}

func MergeSort315(nums [][2]int) {
	l := len(nums)
	if l < 2 {
		return
	}
	// fmt.Println(nums)
	sort315(nums, 0, l-1)
}

func sort315(nums [][2]int, lo, hi int) {
	//单个元素不用排序
	if lo == hi {
		return
	}
	//防止整型溢出
	mid := lo + (hi-lo)/2
	// fmt.Println("mid", lo, hi, mid)
	//排左边的
	sort315(nums, lo, mid)
	// fmt.Println("left", nums[lo:mid])
	//排右边的
	sort315(nums, mid+1, hi)
	//合并左右两边有序为一个有序数组
	merge3(nums, lo, mid, hi)
}

func merge3(nums [][2]int, lo, mid, hi int) {
	//把数据存入临时数组，以便交换数据
	for i := lo; i <= hi; i++ {
		tmp315[i] = nums[i]
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
			nums[i] = tmp315[m]
			m++
		} else if m == hi+1 {
			//右边的数组已合并完, 如果左边还有剩余，全部merger到nums
			// fmt.Println("change2", i, s, tmp[s])
			nums[i] = tmp315[s]
			s++
			count315[nums[i][1]] += m - mid - 1
		} else if tmp315[s][0] > tmp315[m][0] {
			//如果左边的大，把右边小的存入nums对应的位置，右边向往走一格
			// fmt.Println("change3", i, m, tmp[m])
			nums[i] = tmp315[m]
			m++
		} else {
			//如果左边的小，把左边的存入nums对应的位置，左边向前走
			// fmt.Println("change4", i, s, tmp[s])
			nums[i] = tmp315[s]
			s++
			count315[nums[i][1]] += m - mid - 1
		}
	}

	// fmt.Println("nums", nums)
}
