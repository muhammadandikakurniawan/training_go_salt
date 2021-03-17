package tests

import (
	"fmt"
	"testing"

	"github.com/muhammadandikakurniawan/training_go_salt/packages"
)

func TestAddRight(t *testing.T) {
	node := &packages.BinaryTree{}

	node.Insert(32)
	node.Insert(15)
	node.Insert(10)
	node.Insert(23)
	node.Insert(64)
	node.Insert(43)
	node.Insert(74)

	node.Print()
	fmt.Print(node)
}
