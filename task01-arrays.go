package homework

func average(input [15]float32) (result float32) {
	for _, v := range input {
		var sum float32
		sum += v
		result = sum / 4
	}
	return result
}
