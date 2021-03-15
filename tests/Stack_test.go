package tests

import (
	"fmt"
	"testing"

	"github.com/muhammadandikakurniawan/training_go_salt/packages"
)

func TestPushStack(t *testing.T) {
	node := packages.NewStack(10)

	node.Push(30)
	node.Push(50)
	node.Push(70)

	popres := node.Pop()
	popres1 := node.Pop()
	popres2 := node.Pop()
	popres3 := node.Pop()

	fmt.Print(popres)
	fmt.Print(popres1)
	fmt.Print(popres2)
	fmt.Print(popres3)
	fmt.Print(node)
}
