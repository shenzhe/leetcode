package code

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	arr     []int
	len     int
	mapping map[int]int
}

func NewRandomizedSet() RandomizedSet {
	return RandomizedSet{
		arr:     make([]int, 0),
		len:     0,
		mapping: make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	//看mapping里有没有，有直接返回false
	_, ok := rs.mapping[val]
	if ok {
		return false
	}
	rs.arr = append(rs.arr, val)
	rs.mapping[val] = rs.len
	rs.len++
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	//找不到元素，删除
	idx, ok := rs.mapping[val]
	if !ok {
		return false
	}
	//不是最后一个元素，交换到最后一个元素
	rs.len--
	if idx != rs.len {
		lastVal := rs.arr[rs.len]
		lastIdx := rs.mapping[lastVal]
		//交换
		// fmt.Println("remove", idx, val, lastIdx, lastVal)
		rs.arr[idx], rs.arr[lastIdx] = rs.arr[lastIdx], rs.arr[idx]
		rs.mapping[lastVal] = idx
	}
	//删除mapping
	delete(rs.mapping, val)
	//删除arr最后一个元素
	rs.arr = rs.arr[:rs.len]
	return true

}

func (rs *RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(rs.len)
	return rs.arr[randIdx]
}

func (rs *RandomizedSet) GetRandomAndIdx() (int, int) {
	randIdx := rand.Intn(rs.len)
	return rs.arr[randIdx], randIdx
}

func (rs *RandomizedSet) PrintAllMapping() {
	for k, v := range rs.mapping {
		fmt.Println(k, v)
	}
}

func (rs *RandomizedSet) Get(idx int) int {
	return rs.arr[idx]
}

func (rs *RandomizedSet) GetIdx(val int) int {
	idx, ok := rs.mapping[val]
	if ok {
		return idx
	}
	return -1
}

func (rs *RandomizedSet) GetAll() []int {
	return rs.arr
}

func (rs *RandomizedSet) Len() int {
	return rs.len
}
