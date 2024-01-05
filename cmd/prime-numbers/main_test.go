package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIsPrime(t *testing.T) {

	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is prime number"},
		{"not prime", 8, false, "8 is not prime because it is devisable by 2"},
		{"zero", 0, false, "0 by definition is not a prime number"},
		{"Negative number", -4, false, "Negative numbers by definition are not prime"},
	}

	for _, e := range primeTests {

		result, msg := isPrime(e.testNum)

		if e.expected && !result {

			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {

			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s, but got %s", e.name, e.msg, msg)
		}

	}
}

func TestPrompt(t *testing.T) {

	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	prompt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if string(out) != "=>" {
		t.Errorf("incorrect prompt: expcted => but got %s", out)
	}

}

func TestIntro(t *testing.T) {

	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a number") {
		t.Errorf("incorrect intro: expcted string to contain Enter a number, but got %s", out)
	}

}

func TestCheckNumbers(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "please enter a number"},
		{name: "zero", input: "0", expected: "0 by definition is not a prime number"},
		{name: "one", input: "1", expected: "1 by definition is not a prime number"},
		{name: "two", input: "2", expected: "2 is prime number"},
		{name: "three", input: "3", expected: "3 is prime number"},
		{name: "negative", input: "-4", expected: "Negative numbers by definition are not prime"},
		{name: "four", input: "4", expected: "4 is not prime because it is devisable by 2"},
		{name: "typed", input: "sdfdf", expected: "please enter a number"},
		{name: "quit", input: "q", expected: ""},
	}

	for _, e := range tests {

		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)

		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}

func TestReadUserInput(t *testing.T) {

	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)

	<-doneChan

	close(doneChan)

}
