package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneChan := make(chan bool)

	go readUserInput(doneChan)

	<-doneChan

	close(doneChan)

	fmt.Println("Goodbye !")

}

func intro() {

	fmt.Println("Is it Prime number ? ")
	fmt.Println("--------------------")
	fmt.Println("Enter a number, and i will tell you if it is a prime one !, Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("=> ")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {

	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	num, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "please enter a number", false
	}

	_, msg := isPrime(num)

	return msg, false
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
