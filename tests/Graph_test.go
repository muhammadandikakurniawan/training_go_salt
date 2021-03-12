package tests

import (
	"testing"

	"github.com/muhammadandikakurniawan/training_go_salt/packages"
)

func TestNeighBors(t *testing.T) {
	graph := packages.InitGraph()

	graph.AddNode() //0
	graph.AddNode() //1
	graph.AddNode() //2
	graph.AddNode() //3

	graph.AddEdges(0, 1, 10)
	graph.AddEdges(0, 2, 12)
	graph.AddEdges(0, 3, 17)

	graph.AddEdges(1, 2, 15)
	graph.AddEdges(1, 3, 18)

	newighbors := graph.NeighBors(1)

	if len(newighbors) != 3 {
		t.Fatal()
	}

}

func TestGetEdges(t *testing.T) {
	graph := packages.InitGraph()

	graph.AddNode() //0
	graph.AddNode() //1
	graph.AddNode() //2
	graph.AddNode() //3

	graph.AddEdges(0, 1, 10)
	graph.AddEdges(0, 2, 12)
	graph.AddEdges(0, 3, 17)

	graph.AddEdges(1, 2, 15)
	graph.AddEdges(1, 3, 18)

	graph.AddEdges(2, 3, 20)

	edges := graph.GetEdges()

	if len(edges) != 6 {
		t.Fatal()
	}

}

func TestGetNodes(t *testing.T) {
	graph := packages.InitGraph()

	graph.AddNode() //0
	graph.AddNode() //1
	graph.AddNode() //2
	graph.AddNode() //3

	graph.AddEdges(0, 1, 10)
	graph.AddEdges(0, 2, 12)
	graph.AddEdges(0, 3, 17)

	graph.AddEdges(1, 2, 15)
	graph.AddEdges(1, 3, 18)

	graph.AddEdges(2, 3, 20)

	nodes := graph.GetNodes()

	if len(nodes) != 4 {
		t.Fatal()
	}

}

func TestAddNode(t *testing.T) {
	graph := packages.InitGraph()

	graph.AddNode()           //0
	graph.AddNode()           //1
	graph.AddNode()           //2
	addRes := graph.AddNode() //3

	if addRes != 3 {
		t.Fatal()
	}

}
