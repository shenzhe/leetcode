package code

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) UnionFind {
	uf := UnionFind{
		n,
		make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) FindRoot(p int) int {
	// fmt.Printf("p=%d, parent=%d\n", p, uf.parent[p])
	if uf.parent[p] != p {
		// fmt.Printf("not equal, p=%d, parent=%d\n", p, uf.parent[p])
		uf.parent[p] = uf.FindRoot(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UnionFind) Union(p, q int) {
	rp := uf.FindRoot(p)
	rq := uf.FindRoot(q)
	if rp == rq {
		return
	}
	uf.parent[rp] = rq
	uf.count--
}

func (uf UnionFind) Connected(p, q int) bool {
	return uf.FindRoot(p) == uf.FindRoot(q)
}

func (uf UnionFind) GetParent() []int {
	return uf.parent
}
