package homework

func reverse(input []int64) (result []int64) {
	for _, val := range input {
		result = append([]int64{val}, result...)
	}
	return
}
