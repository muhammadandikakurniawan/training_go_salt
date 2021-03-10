package packages

func Fibonacci(idx uint) (int64, []int64) {
	result := []int64{0, 1}

	if idx >= 2 {
		for i := uint(2); i < idx; i++ {
			result = append(result, result[i-1]+result[i-2])
		}
	}

	return result[idx-1], result
}
