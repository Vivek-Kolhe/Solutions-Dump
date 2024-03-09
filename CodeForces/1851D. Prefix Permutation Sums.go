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
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n-1)
		for i := 0; i < n-1; i++ {
			fmt.Fscan(in, &arr[i])
		}

		seen := make(map[int]bool)
		for i := 1; i <= n; i++ {
			seen[i] = false
		}
		sum := (n * (n + 1)) / 2
		if sum < arr[n-2] {
			fmt.Fprintln(out, "NO")
			continue
		}
		var missed int
		if sum > arr[n-2] {
			missed = sum - arr[n-2]
		}
		for i := 1; i < n-1; i++ {
			diff := arr[i] - arr[i-1]
			if seen[diff] {
				missed = diff
			}
			if diff <= n {
				seen[diff] = true
			} else {
				missed = diff
			}
		}
		if arr[0] >= 1 && arr[0] <= n {
			seen[arr[0]] = true
		}
		if missed == 0 {
			missed = arr[0]
		}
		for k, v := range seen {
			if !v {
				missed -= k
			}
		}
		if missed == 0 {
			fmt.Fprintln(out, "YES")
			continue
		}
		fmt.Fprintln(out, "NO")
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
