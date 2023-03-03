package code

func Execuate_990(equations []string) bool {
	return equationsPossible(equations)
}

func equationsPossible(equations []string) bool {

	uf := NewUnionFind(26)

	diff := make([][2]int, 0)
	for _, equation := range equations {
		if equation[1] == '!' {
			diff = append(diff, [2]int{getInt(equation[0]), getInt(equation[3])})
		} else {
			uf.Union(getInt(equation[0]), getInt(equation[3]))
		}
	}

	for _, v := range diff {
		if uf.Connected(v[0], v[1]) {
			return false
		}
	}

	return true
}

func getInt(b byte) int {
	return int(b) - 'a'
}
