package tests

import (
	"fmt"
	"testing"

	"github.com/muhammadandikakurniawan/training_go_salt/packages"
)

func TestArrayset(t *testing.T) {
	arr := []interface{}{1, 4}

	toSet := packages.ArrayToSet(arr)

	fmt.Println(toSet)
}
