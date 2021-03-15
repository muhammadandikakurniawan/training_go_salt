package tests

import (
	"testing"

	"github.com/muhammadandikakurniawan/training_go_salt/packages"
)

func TestInsertBegin(t *testing.T) {
	node := packages.NewLinkedList(100)
	node.InsertBegin(200)

	if node.Value != 200 && node.Next.Value != 100 {
		t.Error()
	}
}
