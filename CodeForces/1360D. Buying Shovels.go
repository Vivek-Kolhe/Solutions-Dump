package main
 
import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
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
 
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
 
func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
 
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(in, &n, &k)
 
		ans := n
		for i := 1; i*i <= n; i++ {
			if n%i == 0 {
				if i <= k {
					ans = Min(ans, n/i)
				}
				if n/i <= k {
					ans = Min(ans, i)
				}
			}
		}
		fmt.Fprintln(out, ans)
	}
}
 
func main() {
	Solve(os.Stdin, os.Stdout)
}
