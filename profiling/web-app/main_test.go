package main

import "testing"

func TestExponentialFibonacci(t *testing.T) {
	expected := 1134903170
	got := exponentialFibonacci(45)
	if got != expected {
		t.Error("expected and got are not equal", expected, got)
	}
}

func BenchmarkExponentialFibonacci(b *testing.B) {
	for n:=0; n < b.N; n++ {
		exponentialFibonacci(25)
	}
}

func BenchmarkLinearFibonacci(b *testing.B) {
	for n:=0; n < b.N; n++ {
		linearFibonacci(25)
	}
}
