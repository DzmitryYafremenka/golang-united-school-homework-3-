package homework

import "sort"

// Task 3: Maps
// function that returns map values sorted in order of increasing keys.
// Input -> {2: "a", 0: "b", 1: "c"}
// Output -> ["b", "c", "a"]
// Input -> {10: "aa", 0: "bb", 500: "cc"}
// Output -> ["bb", "aa", "cc"]

func sortMapValues(input map[int]string) (result []string) {
	var newMap = make([]int, 0)
	var newVal = make([]string, 0)
	for k, _ := range input {
		newMap = append(newMap, k)
	}
	sort.Ints(newMap)
	for _, v := range newMap {
		newVal = append(newVal, input[v])
	}
	return newVal
}
