package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	outArr := []string{}
	keys := make([]int, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	for _, k := range keys {
		outArr = append(outArr, input[k])
	}

	return outArr
}
