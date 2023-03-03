package code

func Execuate_96(n int) int {
	return numTrees(n)
}

var cache_96 map[int]map[int]int

func numTrees(n int) int {
	cache_96 = make(map[int]map[int]int)
	return count(1, n)
}

func count(lo, hi int) int {
	if lo > hi {
		return 1
	}

	if val, ok := cache_96[lo][hi]; ok {
		return val
	}

	res := 0
	for i := lo; i <= hi; i++ {
		left := count(lo, i-1)
		right := count(i+1, hi)
		res += left * right
	}
	if _, ok2 := cache_96[lo]; !ok2 {
		cache_96[lo] = make(map[int]int)
	}
	cache_96[lo][hi] = res
	return res
}
