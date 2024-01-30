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

	var s string
	fmt.Fscan(in, &s)
	
	n := len(s)
	if n < 26 {
		fmt.Fprintln(out, -1)
		return
	}
	idx := make(map[rune]int)
	curr := 'a'
	for i := 0; i < n; i++ {
		if rune(s[i]) <= curr {
			idx[curr] = i
			curr++
		}
		if curr >= 'z' && len(idx) == 26 {
			break
		}
	}
	if len(idx) < 26 {
		fmt.Fprintln(out, -1)
		return
	}
	res := make([]string, n)
	for k, v := range idx {
		res[v] = string(k)
	}
	for i := 0; i < n; i++ {
		if res[i] == "" {
			res[i] = string(s[i])
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprint(out, res[i])
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
