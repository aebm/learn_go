package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberstoSum ...[]int) []int {
	lengthOfNumbers := len(numberstoSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numberstoSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
