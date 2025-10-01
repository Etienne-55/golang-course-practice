package main


func FactorialRecursive(number int) int {
	if number == 0 {
		return 1
	}
	return number * FactorialRecursive(number - 1)
}

func Factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

func Sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

