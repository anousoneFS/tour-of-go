package main

func CountChar(s string) map[string]int {
	result := make(map[string]int)
	step := 0
	// loop over string
	for _, v := range s {
		step = 0
		count := 0
		for key := range result {
			// check if next word(v) already in result
			if key == string(v) {
				step = 1
				break
			}
		}
		if step == 0 {
			for _, val := range s {
				if v == val {
					count++
				}
			}
			result[string(v)] = count
		}
	}
	return result
}
