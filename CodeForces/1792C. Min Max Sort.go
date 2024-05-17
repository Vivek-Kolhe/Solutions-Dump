package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Func(mid, n int, arr []int) bool {
	temp := make([]int, 0)
	for _, v := range arr {
		if v > mid && v < n-mid+1 {
			temp = append(temp, v)
		}
	}

	for i := 1; i < len(temp); i++ {
		if temp[i-1] > temp[i] {
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
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		low, high := 0, n / 2
		ans := n / 2
		for low <= high {
			mid := (low + high) >> 1
			if Func(mid, n, arr) {
				ans = mid
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
