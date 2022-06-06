package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var res float32
	var n int
	for i := 0; i < len(input); i++ {
		if input[i] > 0 {
			res += input[i]
			n++
		}
	}
	return float32(res) / float32(n)
}
