package main

import (
	"strings"
)

// count := len(strings.Fields("sone der anousone haha haha der")) fmt.Println(count)

func WordCount(s string) map[string]int {
	words := strings.Fields(s) // split string to slices
	result := make(map[string]int)
	for _, word := range words {
		result[word] += 1
	}
	return result
}

func WordCountV2(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Split(s, " ")
	// loop over string slices
	for _, words := range words { // index, value in slices
		step := 0
		count := 0
		for key := range result { // key, value in map
			// check if next word not in result
			if words == string(key) {
				step = 1
				break
			}
		}
		if step == 0 {
			for _, v := range words {
				if words == string(v) {
					count++
				}
			}
			result[words] = count
		}
	}
	return result
}
