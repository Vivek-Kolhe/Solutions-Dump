package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Pair struct {
	First  int
	Second int
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k, q int
		fmt.Fscan(in, &n, &k, &q)
		a, b := make([]int, k+1), make([]int, k+1)
		for i := 1; i <= k; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 1; i <= k; i++ {
			fmt.Fscan(in, &b[i])
		}
		for ; q > 0; q-- {
			var l int
			fmt.Fscan(in, &l)

			idx := -1
			low, high := 0, k
			for low <= high {
				mid := (low + high) >> 1
				if a[mid] == l {
					idx = mid
					break
				}
				if a[mid] < l {
					idx = mid
					low = mid + 1
				} else {
					high = mid - 1
				}
			}

			if a[idx] == l {
				fmt.Fprint(out, b[idx], " ")
				continue
			}
			ans := b[idx]
			remDist := l - a[idx]
			dist := a[idx+1] - a[idx]
			time := b[idx+1] - b[idx]
			ans += (remDist * time / dist)
			fmt.Fprint(out, ans, " ")
		}
		fmt.Fprintln(out)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
