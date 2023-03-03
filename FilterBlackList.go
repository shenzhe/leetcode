package code

import "math/rand"

type FilterBlackList struct {
	mapping map[int]int
	sz      int
}

func NewFilterBlackList(n int, blacklist []int) FilterBlackList {
	fbl := FilterBlackList{
		mapping: make(map[int]int),
	}
	fbl.sz = n - len(blacklist)

	/*
		建立黑名单mapping一个map, 后续会把黑名单里的值替换成正确的值
		把黑名单转为白名单
	*/
	for _, v := range blacklist {
		fbl.mapping[v] = -1
	}

	//从后往前替换
	last := n - 1
	/*
		遍历blacklist, 把黑名单对应的值，替换成正确的值
	*/
	for _, v := range blacklist {
		//黑名单的值本身小于sz，才有替换的必要
		if v < fbl.sz {
			for {
				//如果last在黑名单里, 跳转，继续往前走
				if _, ok := fbl.mapping[last]; ok {
					last--
				} else {
					/*
						说明不是black值，放到mapping里，完成转换
						并跳出for, 继续替换下一个black
					*/
					fbl.mapping[v] = last
					break
				}
			}
			//完成替换，往前走一格
			last--
		}
	}

	return fbl
}

func (f *FilterBlackList) Pick() int {
	//随机取一个<sz的数
	idx := rand.Intn(f.sz)
	//如果在mapping里，说明原本是黑名单里的，直接取mapping值
	if v, ok := f.mapping[idx]; ok {
		return v
	}
	return idx
}
