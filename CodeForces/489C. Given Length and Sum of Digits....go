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

	var m, s int
	fmt.Fscan(in, &m, &s)

	if (m > 1 && s < 1) || s > m*9 {
		fmt.Fprintln(out, -1, -1)
		return
	}
	maxi := make([]int, m)
	mini := make([]int, m)
	tempsum := s
	for i := 0; i < m; i++ {
		if tempsum > 9 {
			maxi[i] = 9
			tempsum -= 9
		} else {
			maxi[i] = tempsum
			tempsum = 0
		}
	}
	s -= 1
	for i := m-1; i > 0; i-- {
		if s > 9 {
			mini[i] = 9
			s -= 9
		} else {
			mini[i] = s
			s = 0
		}
	}
	mini[0] = s + 1
	for _, v := range mini {
		fmt.Fprint(out, v)
	}
	fmt.Fprint(out, " ")
	for _, v := range maxi {
		fmt.Fprint(out, v)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
