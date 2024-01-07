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

	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n/2)
	res := make([]int, n)
	for i := 0; i < n/2; i++ {
		fmt.Fscan(in, &arr[i])
	}

	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = arr[i] / 2, arr[i] / 2
		if arr[i] % 2 == 1 {
			res[n-1-i] += 1
		}
	}
	res[0], res[n-1] = 0, arr[0]
	for i := 1; i < n/1; i++ {
		diff := res[i] - res[i-1]
		res[i] = res[i-1]
		res[n-i-1] += diff
		if res[n-i-1] > res[n-i] {
			diff = res[n-i-1] - res[n-i]
			res[n-i-1] = res[n-i]
			res[i] += diff
		}
	}
	for _, v := range res {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
