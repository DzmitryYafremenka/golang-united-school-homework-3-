package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
	var cup int
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			cup++
		}
		sum += input[i]
	}
	return sum / float32(cup)
}
