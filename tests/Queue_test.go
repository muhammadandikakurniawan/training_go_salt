package tests

import (
	"testing"

	"github.com/muhammadandikakurniawan/training_go_salt/packages"
)

func TestPushQueue(t *testing.T) {
	node := packages.NewQueue()

	node.Push(800)
	node.Push(900)
	node.Push(1500)
	node.Push(1800)

	if len(*node) != 4 {
		t.Fatal()
	}
}

func TestDeQueue(t *testing.T) {
	node := packages.NewQueue()

	node.Push(800)
	node.Push(900)
	node.Push(1500)
	node.Push(1800)

	d1 := node.Dequeue()
	d2 := node.Dequeue()

	if len(*node) != 2 && d1 != 800 && d2 != 900 {
		t.Fatal()
	}
}

func TestIsEmptyQueue(t *testing.T) {
	node := packages.NewQueue()

	node.Print()

	node.Push(800)
	node.Push(900)
	node.Push(1500)
	node.Push(1800)

	node.Print()

	d1 := node.Dequeue()
	d2 := node.Dequeue()
	d3 := node.Dequeue()
	d4 := node.Dequeue()
	node.Print()
	if !node.IsEmpty() && d1 != 800 && d2 != 900 && d3 != 1500 && d4 != 1800 {
		t.Fatal()
	}
}
