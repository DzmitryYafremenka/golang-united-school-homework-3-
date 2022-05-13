package homework

func reverse(input []int64) (result []int64) {
	for i:=0; i < len(input)/2; i = i + 1 {
		var j []int64 = input - 1 - i
		input[i], input[j] = input[j], input[i]
	}
	return input
}
