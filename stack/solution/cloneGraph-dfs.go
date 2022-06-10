package solution

import graph "algorithm-go/stack"

func cloneGraph(node *graph.Node) *graph.Node {
	visit := make(map[*graph.Node]*graph.Node, 0)
	visit[nil] = nil
	return helper(node, visit)
}

func helper(node *graph.Node, visit map[*graph.Node]*graph.Node) *graph.Node {
	if node == nil {
		return nil
	}
	stack := make([]*graph.Node, 0)
	stack = append(stack, node)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if v, ok := visit[top]; ok {
			return v
		}

		clone := &graph.Node{
			Val:       top.Val,
			Neighbors: make([]*graph.Node, 0),
		}
		visit[top] = clone
		for _, vNode := range top.Neighbors {
			clone.Neighbors = append(clone.Neighbors, helper(vNode, visit))
		}
	}
	return visit[node]
}
