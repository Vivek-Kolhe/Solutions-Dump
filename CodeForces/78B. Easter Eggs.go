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

	var n int
	fmt.Fscan(in, &n)
	ans := make([]string, 0)
	colors := []string{"R", "O", "Y", "G", "B", "I", "V"}
	for i := 0; i < n/7; i++ {
		for j := 0; j < 7; j++ {
			ans = append(ans, colors[j])
		} 
	}
	for i := 0; i < n%7; i++ {
		ans = append(ans, colors[3+(i%4)])
	}
	fmt.Fprintln(out, strings.Join(ans, ""))
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
