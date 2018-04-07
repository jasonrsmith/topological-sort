package main

import "fmt"

type Graph struct {
	NodeCount     int
	OutgoingEdges [][]int
	IncomingEdges [][]int
}

func NewGraph(nodeCount int) Graph {
	outgoingEdges := make([][]int, nodeCount)
	incomingEdges := make([][]int, nodeCount)
	for i := range outgoingEdges {
		outgoingEdges[i] = make([]int, 0)
		incomingEdges[i] = make([]int, 0)
	}
	return Graph{
		nodeCount,
		outgoingEdges,
		incomingEdges,
	}
}

func (g *Graph) AddPair(a, b int) {
	if a >= g.NodeCount || b >= g.NodeCount {
		panic("Value in pair exceeds node count")
	}
	g.OutgoingEdges[a] = append(g.OutgoingEdges[a], b)
	g.IncomingEdges[b] = append(g.OutgoingEdges[b], a)
}

func (g *Graph) FindNodesWithNoIncomingEdges() []int {
	result := make([]int, 0)
	for i := range g.IncomingEdges {
		if len(g.IncomingEdges[i]) == 0 {
			result = append(result, i)
		}
	}
	return result
}

func (g *Graph) FindKahnTopology() ([]int, error) {
	nodesWithoutIncomingEdges := g.FindNodesWithNoIncomingEdges()
	if len(nodesWithoutIncomingEdges) == g.NodeCount {
		return nodesWithoutIncomingEdges, nil
	}

	result := make([]int, g.NodeCount)
	incomingEdges := CopyEdges(g.IncomingEdges)
	i := 0
	for len(nodesWithoutIncomingEdges) != 0 {
		if i >= len(result) {
			return make([]int, 0), fmt.Errorf("Error finding topology: graph is cyclic")
		}

		m := nodesWithoutIncomingEdges[0]
		nodesWithoutIncomingEdges = nodesWithoutIncomingEdges[1:]
		result[i] = m

		for _, n := range g.OutgoingEdges[m] {
			incomingEdges = RemoveEdge(n, m, incomingEdges)
			if len(incomingEdges[n]) == 0 {
				nodesWithoutIncomingEdges =
					append(nodesWithoutIncomingEdges, n)
			}
		}

		i++
	}

	return result, nil
}

func CopyEdges(edges [][]int) [][]int {
	copy := make([][]int, len(edges))
	for i := range edges {
		copy[i] = make([]int, len(edges[i]))
		for j := range edges[i] {
			copy[i][j] = edges[i][j]
		}
	}
	return copy
}

func RemoveEdge(m, n int, edges [][]int) [][]int {
	for i, node := range edges[m] {
		if n == node {
			edges[m] = append(edges[m][:i], edges[m][i+1:]...)
			return edges
		}
	}
	return edges
}

func main() {
}
