// github.com/Vivek-Kolhe
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func build(idx, low, high int, arr, seg []int) {
	if low == high {
		seg[idx] = arr[low]
		return
	}
	mid := (low + high) >> 1
	build(2*idx, low, mid, arr, seg)
	build(2*idx+1, mid+1, high, arr, seg)
	seg[idx] = seg[2*idx] & seg[2*idx+1]
}

func query(idx, low, high, l, r int, seg []int) int {
	if low > r || high < l {
		return -1
	}
	if l <= low && high <= r {
		return seg[idx]
	}
	mid := (low + high) >> 1
	return query(2*idx, low, mid, l, r, seg) & query(2*idx+1, mid+1, high, l, r, seg)
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
