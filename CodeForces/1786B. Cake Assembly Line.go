package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, w, h, temp int
		fmt.Fscan(in, &n, &w, &h)
		a, b := make([][]int, n), make([][]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &temp)
			a[i] = []int{temp-w, temp+w}
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &temp)
			b[i] = []int{temp-h, temp+h}
		}

		diff := b[0][1] - a[0][1]
		for i := 0; i < n; i++ {
			a[i][0] += diff
			a[i][1] += diff
		}

		e := b[0][0] - a[0][0]
		maxi := 0
		for i := 0; i < n; i++ {
			maxi = max(maxi, b[i][1]-a[i][1])
		}
		if maxi > e {
			fmt.Fprintln(out, "NO")
			continue
		}

		var c int
		for i := 0; i < n; i++ {
			if a[i][0]+maxi <= b[i][0] && a[i][1]+maxi >= b[i][1] {
				c++
			}
		}
		if c == n {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
