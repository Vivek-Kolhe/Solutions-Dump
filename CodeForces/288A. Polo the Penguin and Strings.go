package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	if k > n || (k <= 1 && n > 1) {
		fmt.Fprintln(out, -1)
		return
	}

	if k == 1 && n == 1 {
		fmt.Fprintln(out, "a")
		return
	}

	var ans strings.Builder
	for i := 0; i < n-k+2; i++ {
		if i%2 == 0 {
			ans.WriteByte('a')
		} else {
			ans.WriteByte('b')
		}
	}
	curr := 'c'
	for i := n-k+2; i < n; i++ {
		ans.WriteByte(byte(curr))
		curr++
	}
	fmt.Fprintln(out, ans.String())
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
