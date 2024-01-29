package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	var q int
	fmt.Fscan(in, &q)
	for ; q > 0; q-- {
		var l, r, k int
		fmt.Fscan(in, &l, &r, &k)
		if k%(r-l+1) == 0 {
			continue
		}
		len := r - l + 1
		k = k % (r - l + 1)
		subs := make([]string, len)
		l--
		r--
		for i := l; i <= r; i++ {
			subs[(i+k-l)%len] = string(s[i])
		}
		ss := strings.Join(subs, "")
		s = s[:l] + ss + s[r+1:]
	}
	fmt.Fprintln(out, s)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
