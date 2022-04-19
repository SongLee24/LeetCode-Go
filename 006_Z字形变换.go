package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	// 数据结构，二维数组
	vec := make([]string, numRows)
	idx := 0
	flag := true
	for _, b := range s {
		vec[idx] += string(b)

		if flag {
			idx++
		} else {
			idx--
		}

		if idx == numRows-1 {
			flag = false
		}
		if idx == 0 {
			flag = true
		}
	}

	res := ""
	for _, ls := range vec {
		res += ls
	}
	return res
}
