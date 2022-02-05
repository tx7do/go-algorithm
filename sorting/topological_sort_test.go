package sorting

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
}

func TestTopologicalSort(t *testing.T) {
	{
		g := NewGraph()

		//g.AddNodes(2, 3, 5, 7, 8, 9, 10, 11)

		_ = g.AddEdge(7, 8)
		_ = g.AddEdge(7, 11)

		_ = g.AddEdge(5, 11)

		_ = g.AddEdge(3, 8)
		_ = g.AddEdge(3, 10)

		_ = g.AddEdge(11, 2)
		_ = g.AddEdge(11, 9)
		_ = g.AddEdge(11, 10)

		_ = g.AddEdge(8, 9)

		result, err := g.TopologicalSort()
		if err != nil {
			panic("cycle detected")
		}

		fmt.Println(result)
	}
}
