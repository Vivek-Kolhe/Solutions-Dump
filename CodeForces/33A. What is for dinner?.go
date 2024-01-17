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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		var r, c int
		fmt.Fscan(in, &r, &c)
		
		if _, ok := mp[r]; !ok {
			mp[r] = c
		} else {
			mp[r] = Min(mp[r], c)
		}
	}
	var sum int
	for _, v := range mp {
		sum += v
	}
	fmt.Fprintln(out, Min(k, sum))
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
