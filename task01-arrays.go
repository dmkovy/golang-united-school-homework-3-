package homework

func average(input [15]float32) (result float32) {
	for i, value := range input {
		result += value / float32(i+1)
	}
	return
}
