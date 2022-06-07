package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var preresult []string
	for i := range input {
		preresult = append(preresult, i)
	}
	sort.Strings(preresult)
	for v := range preresult {
		result = append(result, input[v])
	}
	return result
}
