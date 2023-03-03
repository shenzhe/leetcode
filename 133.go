package code

type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func Execuate_133(node *Node133) *Node133 {
	return cloneGraph(node)
}

var visted map[int]*Node133

func cloneGraph(node *Node133) *Node133 {
	visted = make(map[int]*Node133)
	if node == nil {
		return nil
	}
	return traverse_133(node, node.Val)
}

func traverse_133(node *Node133, n int) *Node133 {
	if v, ok := visted[n]; ok {
		return v
	}
	root := &Node133{
		Val: n,
	}
	visted[n] = root
	for _, sub := range node.Neighbors {
		root.Neighbors = append(root.Neighbors, traverse_133(sub, sub.Val))
	}
	return root
}
