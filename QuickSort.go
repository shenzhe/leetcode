package code

import (
	"fmt"
	"math/rand"
)

var count_qs int

func QuickSort(nums []int) []int {
	right := len(nums)
	if right < 2 {
		return nums
	}
	back_nums := make([]int, right)
	// for k, v := range nums {
	// 	back_nums[k] = v
	// }
	copy(back_nums, nums)
	count_qs = 0
	sort_qs(nums, 0, right-1)
	fmt.Println("count_qs", count_qs)

	count_qs = 0
	// fmt.Println(back_nums)
	sort_qs2(back_nums, 0, right-1)
	fmt.Println("count_qs2", count_qs)
	return nums
}
func sort_qs(nums []int, left, right int) {
	if left < right {
		partition := getPartition(nums, left, right)
		sort_qs(nums, left, partition-1)
		sort_qs(nums, partition+1, right)
	}
}

func getPartition(nums []int, left, right int) int {
	i := rand.Intn(right-left+1) + left
	nums[i], nums[right] = nums[right], nums[i]
	pivot := nums[right] //选用最后一个做为基准

	//遍历，和pivot比较，如果<pivot，交换
	j := left - 1
	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			//j往前走一格，然后交换，以达到在j和它左边的是小于基准值
			j++
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		count_qs++
	}
	j++
	//基准位置的值和right交换，以达到j为基准位置，且值为基准值
	if j != right {
		count_qs++
		nums[j], nums[right] = nums[right], nums[j]
	}
	return j
}

func sort_qs2(nums []int, left, right int) {
	// fmt.Println("op", nums)
	if left < right {
		partition := getPartition2(nums, left, right)
		// fmt.Println("pa", partition)
		sort_qs2(nums, left, partition-1)
		sort_qs2(nums, partition+1, right)
	}
}

func getPartition2(nums []int, left, right int) int {
	i := rand.Intn(right-left+1) + left
	//用双针遍历
	nums[i], nums[left] = nums[left], nums[i]
	pivot := nums[left] //选用第一个数基准
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
			count_qs++
		}
		nums[left], nums[right] = nums[right], nums[left]

		for left < right && nums[left] <= pivot {
			count_qs++
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return left
}
