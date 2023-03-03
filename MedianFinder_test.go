package code

import (
	"fmt"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	cases := map[int]float64{
		1: 1,
		2: 1.5,
		3: 2,
	}
	mf := NewFindMedian()
	for i := 1; i <= 3; i++ {
		mf.AddNum(i)
		fmt.Println("add", i)
		result := mf.FindMedian()
		want := cases[i]
		if want != result {
			t.Errorf("want=%f, but=%f", want, result)
		}
	}
}
