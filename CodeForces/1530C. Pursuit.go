package main

import (
	"bufio"
	"fmt"
	"io"
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

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		a, b := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &b[i])
		}

		sort.Slice(a, func(i, j int) bool {
			return a[i] > a[j]
		})
		sort.Slice(b, func(i, j int) bool {
			return b[i] > b[j]
		})

		var ans int
		low, high := 0, int(10e12)
		for low <= high {
			mid := (low + high) >> 1
			newrounds := mid - n
			best := (mid - mid / 4) - newrounds
			var sum, ilya int
			sum += (100 * newrounds)
			for i := 0; i < best; i++ {
				sum += a[i]
			}
			for i := 0; i < Min(n, mid-mid/4); i++ {
				ilya += b[i]
			}
			if sum >= ilya {
				ans = newrounds
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		fmt.Fprintln(out, Max(ans, 0))
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
