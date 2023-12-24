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

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)
	
	if n < 2 {
		for i := 0; i < k; i++ {
			fmt.Fprint(out, s)
		}
		fmt.Fprintln(out)
		return
	}
	pref := ""
	l := 1
	for l < n {
		if strings.HasPrefix(s, s[l:]) {
			pref = s[l:]
			break
		}
		l++
	}
	fmt.Fprint(out, s)
	for i := 1; i < k; i++ {
		if pref != "" {
			fmt.Fprint(out, s[len(pref):])
		} else {
			fmt.Fprint(out, s)
		}
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
