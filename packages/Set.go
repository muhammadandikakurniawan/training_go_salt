package packages

type Set map[interface{}]interface {
}

func ArrayToSet(arr []interface{}) (result []interface{}) {

	tmpMap := make(map[interface{}]bool)
	for _, e := range arr {

		isExist := tmpMap[e]
		if !isExist {
			result = append(result, e)
			tmpMap[e] = true
		}

	}

	return result
}

func arrangeSubElement(from, next, ne int, els *[]interface{}, resArr *[][]interface{}, arrOrigin []interface{}) {

	// attempCount := 0
	// for i := 0; i < len(arrOrigin); i++{

	// }
	lengthOrigin := len(arrOrigin)
	if from+ne == lengthOrigin && next > lengthOrigin {
		return
	}

	if len(*els) == ne {
		*resArr = append(*resArr, *els)
		*els = []interface{}{}
	}

	if (next + 1) == lengthOrigin {
		from += 1
		next = 0
	}

	*els = append(*els, arrOrigin[from+next])
	next += 1

	arrangeSubElement(from, next, ne, els, resArr, arrOrigin)
}

func arrangeSub(currI, ne uint64, resArr *[]interface{}, arrOrigin []interface{}) {

	// attempCount := 0
	// for i := 0; i < len(arrOrigin); i++{

	// }

	// attempCount := 1
	// for uint64(len(*resArr)) <= ne {

	// 	resArr

	// }

}

func PowerSet(arr []interface{}) (result [][]interface{}) {

	tmpArr := ArrayToSet(arr)
	length := len(tmpArr)
	// sub := &[]interface{}{}
	for n := 0; n < length; n++ {

		arrangeSubElement(0, 0, n)

		// if n == 0 {
		// 	result = append(result, []interface{}{})
		// 	continue
		// }

		// for i, _ := range tmpArr {

		// 	to := i + n
		// 	res := to - length
		// 	if res > 0 {

		// 	} else {

		// 	}

		// 	if n > i {
		// 		break
		// 	}

		// }

	}

	// ss := math.Pow(2, float64(len(arr)))

	// n := 1
	// for ix, ex := range arr {

	// 	for iy, ey := range arr {

	// 		sub :=

	// 		n++
	// 	}
	// 	n = 0
	// }

	return result
}

//0, 1, 2, 3, 4
//1, 2, 3, 4, 5

// n = 0
// => {} : 1

// n = 1
// => [1][2][3][4][5] : 5

// n = 2
// => [1,2][1,3][1,4][1,5][2,3][2,4][2,5][3,4][3,5][4,5] : 10

// n = 3
// => [1,2,3][1,2,4][1,2,5][1,3,4][1,3,5][1,4,5][2,3,4][2,3,5][2,4,5][3,4,5] : 10

// n = 4
// => [1,2,3,4][1,2,3,5][1,2,4,5][1,3,4,5][2,3,4,5] = 5

//n = 5
// => [1,2,3,4,5] : 1
