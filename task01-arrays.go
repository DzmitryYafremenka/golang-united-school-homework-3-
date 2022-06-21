package homework

import (
	"fmt"
)

func average(input [15]float32) (result float32) {
	fmt.Scan(input)
	for _, v := range input {
		result += v
	}
	return result / float32(len(input))
	//kdsjgkjhdsjokahgkjldhsafjokghdjlkhfgkjhdsfg
}
