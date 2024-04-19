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
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		var m int
		fmt.Fscan(in, &m)

		seg := make([]int, 4*n)
		build(1, 0, n-1, arr, seg)
		for m > 0 {
			var l, k int
			fmt.Fscan(in, &l, &k)
			m--

			if k > arr[l-1] {
				fmt.Fprint(out, -1, " ")
				continue
			}
			idx := -1
			low, high := l - 1, n - 1
			for low <= high {
				mid := (low + high) >> 1
				if query(1, 0, n-1, l-1, mid, seg) >= k {
					idx = mid
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			fmt.Fprint(out, idx+1, " ")
		}
		fmt.Fprintln(out)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
