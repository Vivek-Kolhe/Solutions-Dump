package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"math"
)


func Sieve(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false
	for p := 2; p <= int(math.Sqrt(float64(n))); p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}
	res := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			res = append(res, i)
		}
	}
	return res
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
