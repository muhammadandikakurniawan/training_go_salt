package packages

func LengthOfLongestSubstring(s string) int {

	sb := []byte(s)

	sMap := make(map[string][]int)
	// sMap["a"] = append(sMap["a"], 1)
	// rsm := len(sMap["a"])

	res := 0
	start := 0
	end := 0

	ct := 0
	for i, e := range sb {

		v := string(e)
		end = i + 1
		if len(sMap[v]) <= 0 {
			sMap[v] = append(sMap[v], i)

		} else {
			if ct == 1 {
				start = i
				ct = 0
			} else {
				start = sMap[v][0] + 1
				sMap[v][0] = i
				ct = 1
			}

		}

		tmpRes := end - start

		if res < tmpRes {
			res = tmpRes
		}

	}

	return res
}
