package sorting

import "fmt"

type NodeType interface{}
type Node map[interface{}]bool
type NodeSlice []NodeType
type NodeMap map[NodeType]Node

type Adjacency map[NodeType]Node

type Graph struct {
	nodes NodeMap
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(NodeMap),
	}
}

func (n Node) addEdge(name NodeType) {
	n[name] = true
}

func (n Node) edges() NodeSlice {
	var keys NodeSlice
	for k := range n {
		keys = append(keys, k)
	}
	return keys
}

func (g *Graph) AddNode(data NodeType) Node {
	n, ok := g.nodes[data]
	if !ok {
		n = make(Node)
		g.nodes[data] = n
	}
	return n
}

func (g *Graph) AddNodes(data ...NodeType) {
	for _, n := range data {
		g.AddNode(n)
	}
}

func (g *Graph) AddEdge(from, to NodeType) error {
	f := g.AddNode(from)
	g.AddNode(to)
	f.addEdge(to)
	return nil
}

// TopologicalSort 拓扑排序
// @see https://en.wikipedia.org/wiki/Topological_sorting
func (g *Graph) TopologicalSort() (NodeSlice, error) {
	results := NodeSlice{}
	visited := Node{}
	for k := range g.nodes {
		err := g.visit(k, &results, visited)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func (g *Graph) visit(data NodeType, results *NodeSlice, visited Node) error {

	visited[data] = true

	n := g.nodes[data]
	for _, edge := range n.edges() {
		if !visited[data] {
			err := g.visit(edge, results, visited)
			if err != nil {
				return err
			}
		}
	}

	*results = append(*results, data)
	fmt.Println(*results, visited)
	return nil
}
