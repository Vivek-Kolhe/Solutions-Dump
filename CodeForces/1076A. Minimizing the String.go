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
	var s string
	fmt.Fscan(in, &s)

	cnt := make(map[byte]int)
	for i := 0; i < n; i++ {
		cnt[s[i]]++
	}
	if len(cnt) == 1 {
		fmt.Fprintln(out, s[:n-1])
		return
	}
	var idx int
	var flag bool
	for i := 1; i < n; i++ {
		if s[i-1] > s[i] {
			idx = i - 1
			flag = true
			break
		}
	}
	if flag {
		fmt.Fprintln(out, s[:idx]+s[idx+1:])
		return
	}
	fmt.Fprintln(out, s[:n-1])
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
