package syntax

// Sum : return sum value of numbers array using range
func Sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers { // '_' -> ignore
		total += number
	}

	return total
}

// Sum2 : return sum value of numbers array
func Sum2(numbers ...int) int {
	var total int = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	return total
}
