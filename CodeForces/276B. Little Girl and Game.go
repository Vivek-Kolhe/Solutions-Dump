package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	mp := make(map[rune]int)
	for _, v := range s {
		mp[v]++
	}

	var even, odd int
	for _, v := range mp {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd == 0 {
		fmt.Fprintln(out, "First")
		return
	}
	odd--
	if odd%2 == 1 {
		fmt.Fprintln(out, "Second")
		return
	}
	fmt.Fprintln(out, "First")
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
