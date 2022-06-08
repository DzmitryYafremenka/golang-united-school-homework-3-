package homework

import "sort"

type Inputing []int64

func (a Inputing) Len() int {
	return len(a)
}
func (a Inputing) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Inputing) Less(i, j int) bool {
	return a[i] < a[j]
}

func reverse(input []int64) (result []int64) {
	copyinput := make([]int64, 5)
	copy(copyinput, input)
	sort.Sort(sort.Reverse(Inputing(copyinput)))
	result = copyinput
	return result
}
