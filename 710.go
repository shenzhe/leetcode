package code

import "math/rand"

type Solution struct {
	cache map[int]int
	len   int
}

func New710(n int, blacklist []int) Solution {
	s := Solution{
		cache: make(map[int]int),
	}
	l := len(blacklist)
	//真正的长度，也就是随机的最大值
	s.len = n - l
	blackmap := make(map[int]int)
	//建立一个black的map
	for _, v := range blacklist {
		blackmap[v] = 1
	}
	//遍历s.len后面的值，把黑名单洗成白名单
	for i := s.len; i < n; i++ {
		//如果在黑名单中，跳过
		if _, ok := blackmap[i]; ok {
			continue
		}
		for l > 0 {
			l--
			//小于s.len，把black对应的索引值设置为i
			//for是为了跳过大于s.len的黑名单的值
			if blacklist[l] < s.len {
				// fmt.Println("cache", l, blacklist[l], i)
				s.cache[blacklist[l]] = i
				break
			}
		}
		//小优化，如果l=0,说明黑名单已经洗完，可以直接跳出
		if l == 0 {
			break
		}
	}
	// fmt.Println(s.len)
	return s
}

func (s *Solution) Pick() int {
	randIdx := rand.Intn(s.len)
	if v, ok := s.cache[randIdx]; ok {
		return v
	}
	return randIdx
}
