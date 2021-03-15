package tests

import (
	"fmt"
	"testing"

	"github.com/muhammadandikakurniawan/training_go_salt/packages"
)

func TestAddRight(t *testing.T) {
	node := packages.NewBinaryTree(10)

	node.AddRight(30)
	node.AddRight(50)
	node.AddRight(70)

	fmt.Print(node)
}

func TestLeft(t *testing.T) {
	node := packages.NewBinaryTree(10)

	node.AddLeft(30)
	node.AddLeft(50)
	node.AddLeft(70)

	fmt.Print(node)
}

func TestPrint(t *testing.T) {
	node := packages.NewBinaryTree(10)

	node.AddLeft(30)
	node.AddLeft(50)
	node.AddLeft(70)

	fmt.Print(node)
}
