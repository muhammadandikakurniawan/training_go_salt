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

func BenchmarkFibonacci(b *testing.B) {

	for i := 0; i < b.N; i++ {
		res, _ := packages.Fibonacci(6)

		if res != 5 {
			b.Error("salah")
		}
	}

}
