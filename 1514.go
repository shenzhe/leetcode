package code

func Execuate_1514(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	return maxProbability(n, edges, succProb, start, end)
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][][]float64, n)

	for i, v := range edges {
		graph[v[0]] = append(graph[v[0]], []float64{
			float64(v[1]), succProb[i],
		})
		graph[v[1]] = append(graph[v[1]], []float64{
			float64(v[0]), succProb[i],
		})
	}
	dk := NewDijkstraFloat(graph, n)
	dk.SetWeightFunc(func(i1, i2 any) any {
		return i1.(float64) * i2.(float64)
	})
	ans := dk.RunFloat(start, end)
	if ans[end] == -1 {
		return 0
	}
	return ans[end]
}
