package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var preresult []int
	for i := range input {
		preresult = append(preresult, i)
	}
	preresult = sort.IntSlice(preresult)
	for v := range preresult {
		result = append(result, input[v])
	}
	return result
}
