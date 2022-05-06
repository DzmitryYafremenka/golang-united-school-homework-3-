package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	values := []string{}
	keys := []int{}
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		values = append(values, input[k])
	}
	return values
}
