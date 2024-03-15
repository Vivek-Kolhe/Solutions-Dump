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

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, s int
		fmt.Fscan(in, &n, &s)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		pref := make([]int, n)
		pref[0] = arr[0]
		for i := 1; i < n; i++ {
			pref[i] = arr[i] + pref[i-1]
		}
		var curr, maxi, idx int
		for i := 0; i < n; i++ {
			low, high := 0, n-1
			for low <= high {
				mid := (low + high) >> 1
				curr = pref[mid]
				if mid >= i {
					curr -= arr[i]
				}
				if curr > s {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
			if high > maxi {
				maxi = high
				idx = i + 1
			}
		}
		low, high := 0, n - 1
		for low <= high {
			mid := (low + high) >> 1
			if pref[mid] > s {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		if high == n-1 {
			idx = 0
		}
		fmt.Fprintln(out, idx)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
