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

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	o, c := make([]int, 0), make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == ')' {
			c = append(c, i)
		} else {
			o = append(o, i)
		}
	}
	if len(o) != len(c) {
		fmt.Fprintln(out, -1)
		return
	}

	var i, ans int
	for i < len(c) {
		if c[i] < o[i] {
			ans += 2
		}
		i++
	}
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
