package main

import "fmt"

func main() {

	n := 7

	_, msg := isPrime(n)

	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {

	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d by definition is not a prime number", n)
	}

	if n < 0 {
		return false, "Negative numbers by definition are not prime"
	}

	for i := 2; i < n/2; i++ {

		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime because it is devisable by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is prime number", n)
}
