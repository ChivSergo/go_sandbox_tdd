package arrays

func Sum(arrs interface{}) int {
	result := 0
	for _, arr := range arrs {
		result += arr
	}
	return result
}
