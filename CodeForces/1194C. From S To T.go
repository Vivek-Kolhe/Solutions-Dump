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

	var tc int
	fmt.Fscan(in, &tc)
	for ; tc > 0; tc-- {
		var s, t, p string
		fmt.Fscan(in, &s)
		fmt.Fscan(in, &t)
		fmt.Fscan(in, &p)

		var sl, tl int
		tcnt, scnt := make(map[rune]int), make(map[rune]int)
		for _, v := range s {
			sl++
			scnt[v]++
		}
		for _, v := range t {
			tl++
			tcnt[v]++
		}
		for _, v := range p {
			scnt[v]++
		}
		if sl > tl {
			fmt.Fprintln(out, "NO")
			continue
		}
		var i, j int
		for i < sl && j < tl {
			if t[j] == s[i] {
				i++
			}
			j++
		}
		if i != len(s) {
			fmt.Fprintln(out, "NO")
			continue
		}
		flag := true
		for _, v := range t {
			if scnt[v] < tcnt[v] {
				fmt.Fprintln(out, "NO")
				flag = false
				break
			}
		}
		if flag {
			fmt.Fprintln(out, "YES")
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
