package tests

import (
	"testing"

	"github.com/muhammadandikakurniawan/training_go_salt/packages"
)

func TestFibonacci(t *testing.T) {
	res, _ := packages.Fibonacci(6)

	if res != 5 {
		t.Error("salah")
	}
}

func BenchmarkFibonacci(t *testing.B) {
	res, _ := packages.Fibonacci(6)

	if res != 5 {
		t.Error("salah")
	}
}
