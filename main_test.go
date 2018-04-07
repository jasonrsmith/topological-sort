package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	testNodeCount := 4
	g := NewGraph(testNodeCount)
	assert.Equal(t, testNodeCount, len(g.OutgoingEdges))
	assert.Equal(t, 0, len(g.OutgoingEdges[0]))

}

func TestAddPair(t *testing.T) {
	testNodeCount := 4
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	g.AddPair(2, 3)
	assert.Equal(t, 1, g.OutgoingEdges[0][0])
	assert.Equal(t, 3, g.OutgoingEdges[2][0])
}

func TestAddPairPanicsWhenAddingPairExceedsNodeCount(t *testing.T) {
	testNodeCount := 4
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	assert.Panics(t, func() { g.AddPair(4, 1) })
}

func TestFindNodesWithNoIncomingEdgesNoConnections(t *testing.T) {
	testNodeCount := 3
	expectedNodes := []int{0, 1, 2}
	g := NewGraph(testNodeCount)
	nodeList := g.FindNodesWithNoIncomingEdges()
	assert.Equal(t, expectedNodes, nodeList)
}

func TestFindNodesWithNoIncomingEdgesOneEdge(t *testing.T) {
	testNodeCount := 3
	expectedNodes := []int{0, 2}
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	nodeList := g.FindNodesWithNoIncomingEdges()
	assert.Equal(t, expectedNodes, nodeList)
}

func TestFindNodesWithNoIncomingEdgesAllConnected(t *testing.T) {
	testNodeCount := 3
	expectedNodes := []int{}
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	g.AddPair(1, 2)
	g.AddPair(2, 0)
	nodeList := g.FindNodesWithNoIncomingEdges()
	assert.Equal(t, expectedNodes, nodeList)
}

func TestFindKahnTopologyNoEdges(t *testing.T) {
	testNodeCount := 3
	expectedNodes := []int{0, 1, 2}
	g := NewGraph(testNodeCount)
	nodeList, err := g.FindKahnTopology()
	assert.Equal(t, expectedNodes, nodeList)
	assert.Nil(t, err)
}

func TestFindKahnTopologyWithOneEdge(t *testing.T) {
	testNodeCount := 3
	expectedNodes := []int{0, 2, 1}
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	nodeList, err := g.FindKahnTopology()
	assert.Equal(t, expectedNodes, nodeList)
	assert.Nil(t, err)
}

func TestFindKahnTopologyComplex(t *testing.T) {
	testNodeCount := 8
	expectedNodes := []int{0, 1, 2, 3, 4, 5, 7, 6}
	g := NewGraph(testNodeCount)
	g.AddPair(0, 3)
	g.AddPair(1, 3)
	g.AddPair(1, 4)
	g.AddPair(2, 4)
	g.AddPair(2, 7)
	g.AddPair(3, 5)
	g.AddPair(3, 6)
	g.AddPair(3, 7)
	g.AddPair(4, 6)

	nodeList, err := g.FindKahnTopology()
	assert.Equal(t, expectedNodes, nodeList)
	assert.Nil(t, err)
}

func TestCopyEdges(t *testing.T) {
	edges := [][]int{
		{0, 1},
		{1, 2},
	}
	assert.Equal(t, edges, CopyEdges(edges))
}

func TestRemoveEdge(t *testing.T) {
	edges := [][]int{{1, 2, 3, 4}}
	expectedResult := [][]int{{1, 3, 4}}
	actualResult := RemoveEdge(0, 2, edges)
	assert.Equal(t, expectedResult, actualResult)
}
