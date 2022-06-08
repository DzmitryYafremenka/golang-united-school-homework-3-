package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var N float32
	for _, v := range input {
		if v > 0 {
			sum += v
			N++
		} else {
			break
		}
	}
	result = sum / N
	return result
}
