package main

import "testing"

func TestPrimalityComboCheck(t *testing.T) {
	sieve, _ := sieveOfEratosthenes(8000, 8000)
	if !checkPrimalityOfCombo(7, 3, sieve) {
		t.Error("7 and 3 should produce primes when combined")
	}
	if !checkPrimalityOfCombo(7, 109, sieve) {
		t.Error("7 and 3 should produce primes when combined")
	}
	if checkPrimalityOfCombo(13, 5, sieve) {
		t.Error("13 and 5 should not produce primes when combined")
	}
	if checkPrimalityOfCombo(5, 13, sieve) {
		t.Error("13 and 5 should not produce primes when combined")
	}
}
