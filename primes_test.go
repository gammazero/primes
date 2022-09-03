package main

import "testing"
import "fmt"

func Test7919(t *testing.T) {
	p := primes(7919)
	if len(p) != 1000 {
		t.Fatal("did not find 100 primes")
	}
}

func TestCount(t *testing.T) {
	p := primes(100)
	counts := []int{4, 25, 168, 1229, 9592, 78498, 664579}
	x := int64(1)
	for i := range counts {
		x *= 10
		p = primes(x)
		fmt.Println("Found", len(p), "<=", x)
		if len(p) != counts[i] {
			t.Fatal("expected ", counts[i], "primes <= ", x)
		}
	}
}

func BenchmarkPrimes50k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes(50000)
	}
}

func BenchmarkPrimes100k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes(100000)
	}
}

func BenchmarkPrimes500k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes(500000)
	}
}

func BenchmarkPrimes1M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes(1000000)
	}
}
