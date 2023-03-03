package code

import "math/rand"

func Execuate_215(nums []int, k int) int {
	return findKthLargest(nums, k)
}

func findKthLargest(nums []int, k int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	//left = l - k - 1
	//right = k
	// fmt.Println(nums)
	return sort_215(nums, 0, l-1, l-k)
}

func sort_215(nums []int, lo, hi, k int) int {

	if lo < hi {
		partition := partition_215(nums, lo, hi)
		// fmt.Println(lo, hi, k, partition, nums)
		if partition == k {
			return nums[partition]
		} else if partition > k {
			return sort_215(nums, lo, partition-1, k)
		} else {
			return sort_215(nums, partition+1, hi, k)
		}
	}
	return nums[k]
}

func partition_215(nums []int, lo, hi int) int {
	//取第一个做为支点，看他应该在数组的位置
	//随机取一个值放到最后,期望到O(N)的时间复杂度
	i := rand.Intn(hi-lo+1) + lo
	nums[i], nums[hi] = nums[hi], nums[i]
	pivot := nums[hi]
	j := lo - 1
	for i := lo; i < hi; i++ {
		if nums[i] < pivot {
			j++
			//交换
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	j++
	if j != hi {
		nums[j], nums[hi] = nums[hi], nums[j]
	}

	return j
}
