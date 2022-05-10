package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var count int
	for i:= 0; i< len(input); i++ {
		// fmt.Println(input[i])
		if input[i]== 0{
			continue
		}
		count++
		sum += input[i]
	}
	return sum / float32(count)
}
