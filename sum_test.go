package main

import "testing"

func TesteSum(t *testing.T) {
	result := sum(2, 3)
	if result != 5 {
		t.Error("the result must be five")
	}
}
