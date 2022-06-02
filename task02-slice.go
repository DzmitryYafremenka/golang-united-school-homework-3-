package homework

import "sort"

func reverse(input []int64) (result []int64) {
	sort.IntSlice
	for i := range input {
		result = append(result, input[len(input)-1-i])
	}
	return result
}
