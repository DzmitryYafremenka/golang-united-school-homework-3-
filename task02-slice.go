package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	for i := len(input) - 1; i >= len(input)/2; i-- {
		b := len(input) - 1 - i
		input[i], input[b] = input[b], input[i]
	}
	return input
}
