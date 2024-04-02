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
	stack := make([]string, 0)
	for ; n > 0; n-- {
		var s string
		fmt.Fscan(in, &s)
		if s == "cd" {
			var t string
			fmt.Fscan(in, &t)
			s += " " + t
		}

		if s == "pwd" {
			ans := strings.Join(stack, "/")
			if len(ans) < 1 {
				fmt.Fprintln(out, "/")
			} else {
				fmt.Fprintf(out, "/%s/\n", ans)
			}
			continue
		}
		cmd := strings.Split(s, " ")[1]
		if cmd[0] == '/' {
			cmd = cmd[1:]
			stack = []string{}
		}
		path := strings.Split(cmd, "/")
		for _, p := range path {
			if p != ".." {
				stack = append(stack, p)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
