package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func doDamage(n, k int, arr []int) int {
	dmg := 0
	for i := 0; i < n-1; i++ {
		if arr[i+1]-arr[i] > k {
			dmg += k
		} else {
			dmg += arr[i+1] - arr[i]
		}
	}
	dmg += k
	return dmg
}

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, h int
		fmt.Fscan(in, &n, &h)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		low := 1
		high := int(math.Pow(10, 18))
		for low <= high {
			var mid int
			mid = (low + high) / 2
			if doDamage(n, mid, arr) >= h {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		fmt.Fprintln(out, low)
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
