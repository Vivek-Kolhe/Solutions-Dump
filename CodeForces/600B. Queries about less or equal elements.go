package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	a, b := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	sort.Ints(a)
	res := make([]int, m)
	for i, v := range b {
		low, high := 0, n - 1
		for low <= high {
			mid := (low + high) >> 1
			if a[mid] <= v {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		res[i] = low
	}
	for _, v := range res {
		fmt.Fprint(out, v, " ")
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
