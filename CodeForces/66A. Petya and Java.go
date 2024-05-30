package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n string
	fmt.Fscan(in, &n)

	if (n[0] == '-' && len(n) > 20) || len(n) > 19 {
		fmt.Fprintln(out, "BigInteger")
		return
	}

	num, err := strconv.Atoi(n)
	if num >= -128 && num <= 127 {
		fmt.Fprintln(out, "byte")
	} else if num >= -32768 && num <= 32767 {
		fmt.Fprintln(out, "short")
	} else if num >= -2147483648 && num <= 2147483647 {
		fmt.Fprintln(out, "int")
	} else if err == nil {
		fmt.Fprintln(out, "long")
	} else {
		fmt.Fprintln(out, "BigInteger")
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
