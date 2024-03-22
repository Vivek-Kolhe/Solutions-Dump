package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type Pair struct {
	First  int64
	Second int64
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

	var n int64
	fmt.Fscan(in, &n)
	pts := make([]Pair, n)
	for i := int64(0); i < n; i++ {
		var x, y int64
		fmt.Fscan(in, &x, &y)
		p := Pair{
			First: x,
			Second: y,
		}
		pts[i] = p
	}

	sort.Slice(pts, func(i, j int) bool {
		return pts[i].First < pts[j].First
	})
	mini, maxi := pts[0], pts[n-1]
	xdist := Abs(mini.First-maxi.First)
	sort.Slice(pts, func(i, j int) bool {
		return pts[i].Second < pts[j].Second
	})
	mini, maxi = pts[0], pts[n-1]
	ydist := Abs(mini.Second-maxi.Second)
	fmt.Fprintln(out, Max(xdist, ydist)*Max(xdist, ydist))
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
