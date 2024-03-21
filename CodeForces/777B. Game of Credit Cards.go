package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Pair struct {
	First  int
	Second int
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](nums ...T) T {
	res := nums[0]
	for _, v := range nums {
		if res > v {
			res = v
		}
	}
	return res
}

func Max[T Number](nums ...T) T {
	res := nums[0]
	for _, v := range nums {
		if res < v {
			res = v
		}
	}
	return res
}

func Abs[T Number](num T) T {
	if num < 0 {
		return num * -1
	}
	return num
}

func GCD(a, b int) int {
	for a > 0 && b > 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	if a == 0 {
		return b
	}
	return a
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	var a, b string
	fmt.Fscan(in, &a)
	fmt.Fscan(in, &b)

	sher, mor := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		sher[i], _ = strconv.Atoi(string(a[i]))
	}
	for i := 0; i < n; i++ {
		mor[i], _ = strconv.Atoi(string(b[i]))
	}

	sort.Ints(sher)
	sort.Ints(mor)
	var mini, maxi, l, r int
	for l < n && r < n {
		for r < n && sher[l] > mor[r] {
			r++
		}
		if r < n {
			mini++
			l++
			r++
		}
	}
	l, r = 0, 0
	for l < n && r < n {
		for r < n && sher[l] >= mor[r] {
			r++
		}
		if r < n {
			maxi++
			l++
			r++
		}
	}
	fmt.Fprintf(out, "%d\n%d", n-mini, maxi)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
