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

func Check(n, m, mid int, city, tower []int) bool {
	var i, j, c int

	for i < n && j < m {
		if tower[j]-mid <= city[i] && city[i] <= tower[j]+mid {
			i++; c++;
		} else {
			j++
		}
	}
	return c == n
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	city, tower := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &city[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &tower[i])
	}

	low, high := 0, int(1e10)
	for low <= high {
		mid := (low + high) >> 1
		if Check(n, m, mid, city, tower) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Fprintln(out, low)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
