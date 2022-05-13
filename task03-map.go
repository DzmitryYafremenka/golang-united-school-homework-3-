package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	
	values := []string{}
	keys := make([]int, 0, len(input))

	for k := range(input) {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range(keys) {
		values = append(values, input[k])
	}

	return values
}
