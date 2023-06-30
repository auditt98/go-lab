package main

import (
	"fmt"
	"math"
	"strconv"
)

func validate(n string) int {
	num, err := strconv.Atoi(n)
	if err != nil {
		panic("Invalid number")
	}
	return num
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func listPrimes(n int) []int {
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {

	var inp string
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&inp)
	if err != nil {
		panic("Error reading input")
	}
	num := validate(inp)
	for _, v := range listPrimes(num) {
		fmt.Printf("%d ", v)
	}
}
