package code

import "fmt"

type NumArray struct {
	pre []int
}

func NewNumArray(nums []int) NumArray {
	len := len(nums)
	a := NumArray{}
	a.pre = make([]int, len+1)
	for i := 1; i <= len; i++ {
		a.pre[i] = a.pre[i-1] + nums[i-1]
	}
	return a
}

func (na *NumArray) sumRange(i, j int) int {
	return na.pre[j+1] - na.pre[i]
}

func main() {
	fmt.Printf("hellow")
}
