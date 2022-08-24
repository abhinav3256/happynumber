package main

import "fmt"

func main() {
	n := 181
	res := isHappy(n)
	fmt.Println(res)
}

func isHappy(n int) bool {
	sum := squareSum(n)
	fmt.Println(sum)
	if sum == 1 || sum == 7 {
		return true
	}
	return false
}

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		rem := n % 10
		sum = sum + (rem * rem)
		n = n / 10

	}

	if sum > 9 {
		sum = squareSum(sum)
	}
	return sum

}
