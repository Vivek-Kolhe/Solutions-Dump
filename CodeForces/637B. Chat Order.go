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
	var name string
	fmt.Fscan(in, &n)
	stack := make([]string, 0)
	uniq := make(map[string]bool)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &name)
		if _, ok := uniq[name]; !ok {
			uniq[name] = true
		}
		stack = append(stack, name)
	}
	seen := make(map[string]bool)
	for i := len(stack)-1; i >= 0; i-- {
		if len(seen) == len(uniq) {
			break
		}
		if _, ok := seen[stack[i]]; !ok {
			seen[stack[i]] = true
			fmt.Fprintln(out, stack[i])
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
