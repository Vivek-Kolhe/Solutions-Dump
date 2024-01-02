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

	var vh, va, hh, ha, c int
	fmt.Fscan(in, &hh, &ha, &c)
	fmt.Fscan(in, &vh, &va)

	if ha >= vh {
		fmt.Fprintln(out, 1)
		fmt.Fprintln(out, "STRIKE")
		return
	}
	var ans int
	moves := make([]string, 0)
	for vh > 0 {
		ans++
		if vh <= ha {
			vh -= ha
			moves = append(moves, "STRIKE")
			continue
		}
		if hh > va {
			vh -= ha
			moves = append(moves, "STRIKE")
		} else {
			hh += c
			moves = append(moves, "HEAL")
		}
		hh -= va
	}
	fmt.Fprintln(out, ans)
	for i := 0; i < ans; i++ {
		fmt.Fprintln(out, moves[i])
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
