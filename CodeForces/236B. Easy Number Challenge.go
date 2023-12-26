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

	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)

	mod := 1073741824
	facts := make([]int, a*b*c+1)
	for i := 1; i < a*b*c+1; i++ {
		facts[i] = 1
	}
	for i := 2; i < a*b*c+1; i++ {
		for j := i; j < a*b*c+1; j += i {
			facts[j]++
		}
	}
	var ans int
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			for k := 1; k <= c; k++ {
				ans += (facts[i*j*k] % mod)
			}
		}
	}
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
