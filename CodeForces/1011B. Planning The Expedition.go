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

func Assign(cnt map[int]int, mid int) int {
	var c int
	for _, v := range cnt {
		c += (v / mid)
	}
	return c
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &arr[i])
	}

	cnt := make(map[int]int)
	for _, v := range arr {
		cnt[v]++
	}
	low, high := 1, int(10e3)
	for low <= high {
		mid := (low + high) >> 1
		if Assign(cnt, mid) < n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	high = Max(high, 0)
	fmt.Fprintln(out, high)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
