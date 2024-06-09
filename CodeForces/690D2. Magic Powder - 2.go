package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Func(mid, k, n int, a, b []int) bool {
	for i := 0; i < n; i++ {
		need := a[i] * mid
		if b[i] < need {
			if b[i]+k >= need {
				k -= (need - b[i])
			} else {
				return false
			}
		}
	}
	if k < 0 {
		return false
	}
	return true
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	low, high := 0, int(1e10)
	for low <= high {
		mid := (low + high) >> 1
		if Func(mid, k, n, a, b) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Fprintln(out, high)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
