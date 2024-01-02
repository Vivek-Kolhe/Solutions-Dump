package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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

func Dist(a []int) float64 {
	x, y := float64(a[0]), float64(a[1])
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, pop int
	var x, y, k int
	fmt.Fscan(in, &n, &pop)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y, &k)
		arr[i] = []int{x, y, k}
	}

	sort.Slice(arr, func(i, j int) bool {
		return Dist(arr[i]) < Dist(arr[j])
	})
	if pop == 10e5 {
		fmt.Fprintln(out, 0.0)
		return
	}
	for i := 0; i < n; i++ {
		pop += arr[i][2]
		if pop >= 10e5 {
			fmt.Fprintln(out, Dist(arr[i]))
			return
		}
	}
	fmt.Fprintln(out, -1)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
