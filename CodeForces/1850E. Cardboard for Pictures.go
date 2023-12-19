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
		var n, c int64
		fmt.Fscan(in, &n, &c)
		arr := make([]int64, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		var area int64
		low, high := int64(0), int64(10e9)
		for low <= high {
			mid := (low + high) >> 1
			area = 0
			for _, v := range arr {
				area += int64((v + (2 * mid)) * (v + (2 * mid)))
				if area > c {
					break
				}
			}
			if area == c {
				fmt.Fprintln(out, mid)
				break
			} else if area > c {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
