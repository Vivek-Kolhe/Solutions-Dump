package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func build(idx, low, high int, m int64, arr, seg []int64) {
	if low == high {
		seg[idx] = arr[low] % m
		return
	}
	mid := (low + high) >> 1
	build(2*idx, low, mid, m, arr, seg)
	build(2*idx+1, mid+1, high, m, arr, seg)
	seg[idx] = (seg[2*idx] * seg[2*idx+1]) % m
}

func query(idx, low, high, l, r int, m int64, arr, seg []int64) int64 {
	if low > r || high < l {
		return 1
	}
	if l <= low && high <= r {
		return seg[idx]
	}
	mid := (low + high) >> 1
	return (query(2*idx, low, mid, l, r, m, arr, seg) * query(2*idx+1, mid+1, high, l, r, m, arr, seg)) % m
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		var m int64
		var s string
		fmt.Fscan(in, &n, &m)
		arr := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		fmt.Fscan(in, &s)

		seg := make([]int64, 4*n)
		build(1, 0, n-1, m, arr, seg)
		l, r := 0, n - 1
		fmt.Fprint(out, query(1, 0, n-1, l, r, m, arr, seg), " ")
		for i := 0; i < n-1; i++ {
			if s[i] == 'L' {
				l++
			} else {
				r--
			}
			fmt.Fprint(out, query(1, 0, n-1, l, r, m, arr, seg), " ")
		}
		fmt.Fprintln(out)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
