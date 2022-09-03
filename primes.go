package main

// See: https://play.golang.org/p/wcDn6dBb2-F

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func main() {
	var (
		max int64
		pal bool
	)
	flag.Int64Var(&max, "max", 1000, "maximum prime to examine")
	flag.BoolVar(&pal, "pal", false, "composites of palendromic primes")
	flag.Parse()
	if pal {
		palindromes(max)
		return
	}
	for _, p := range primes(max) {
		fmt.Printf("%d ", p)
	}
	fmt.Println()
}

// primes
func primes(n int64) []int64 {
	if n < 2 {
		return nil
	}
	memSize := (n - 1) / 2
	numbers := make([]byte, memSize+1)
	primes := make([]int64, 0, memSize/5)
	primes = append(primes, 2)
	sqn := int64(math.Sqrt(float64(n)))

	for i := int64(3); i <= n; i += 2 {
		// Check if number, without lowest bit, is a multiple of a previously
		// seen prime number.
		if numbers[i>>1] == 1 {
			continue
		}
		primes = append(primes, i)
		if i > sqn {
			continue
		}
		// Mark remaining multiples of i, above i*i, as not prime. Ignore
		// lowest bit since all numbers to test are odd.
		for j := i * i; j <= n; j += (i << 1) {
			numbers[j>>1] = 1
		}
	}
	return primes
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindromes(n int64) {
	plist := primes(n)
	for _, p := range plist {
		pr, _ := strconv.ParseInt(reverse(fmt.Sprintf("%d", p)), 10, 64)
		if pr > p {
			for i := range plist {
				if pr == plist[i] {
					fmt.Println(p, "*", pr, "=", p*pr)
					break
				}
			}
		}
	}
}
