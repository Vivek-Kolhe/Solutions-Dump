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

	var t int
	fmt.Fscan(in, &t)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
