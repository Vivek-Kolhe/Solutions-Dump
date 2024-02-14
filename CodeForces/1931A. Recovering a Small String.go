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

	letters := "aabcdefghijklmnopqrstuvwxyz"
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		var ans string
		if n < 28 {
			ans = "aa" + string(letters[n-2])
			fmt.Fprintln(out, ans)
		} else if n < 53 {
			ans = "a" + string(letters[n-27])  + "z"
			fmt.Fprintln(out, ans)
		} else {
			ans = string(letters[n-52]) + "zz"
			fmt.Fprintln(out, ans)
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
