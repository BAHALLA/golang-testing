package main

import (
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
