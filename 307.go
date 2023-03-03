package code

type Difference struct {
	diff []int
	len  int
}

func NewDifference(nums []int) Difference {
	a := Difference{}
	a.diff = nums
	l := len(nums)
	a.len = l
	for i := 1; i < l; i++ {
		a.diff[i] = nums[i] - nums[i-1]
	}
	return a
}

func (this *Difference) update(s, e, v int) {
	this.diff[s] += v
	e++
	if e < this.len {
		this.diff[e] -= v
	}
}

func (this *Difference) recover() []int {
	r := make([]int, this.len)
	r[0] = this.diff[0]
	for i := 1; i < this.len; i++ {
		r[i] = this.diff[i] + r[i-1]
	}
	return r
}
