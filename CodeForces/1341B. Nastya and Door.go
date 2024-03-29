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
		var n, k int
		fmt.Fscan(in, &n, &k)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		peaks := make([]int, n)
		for i := 1; i < n-1; i++ {
			if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
				peaks[i] = 1
			}
			peaks[i] += peaks[i-1]
		}
		var ans, maxi int
		for i := k-1; i < n; i++ {
			curr := peaks[i-1] - peaks[i-k+1]
			if curr > maxi {
				maxi = curr
				ans = i-k+1
			}
		}
		fmt.Fprintln(out, maxi+1, ans+1)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
