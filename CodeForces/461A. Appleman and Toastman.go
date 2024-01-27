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

	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	if n == 1 {
		fmt.Fprintln(out, arr[0])
		return
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	for i := 1; i < n; i++ {
		arr[i] += arr[i-1]
	}
	ans := arr[n-1]
	for i := n-1; i >= 0; i-- {
		if i == 1 {
			break
		}
		ans += arr[i]
	}
	ans += arr[0] + (arr[1] - arr[0])
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
