package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var m int
	bits := make([]int, 30)
	fmt.Fscan(in, &m)
	for ; m > 0; m-- {
		var op, x int
		fmt.Fscan(in, &op, &x)

		c := 0
		if op == 1 {
			bits[x]++
			c++
		} else {
			if c < 1 && x == 0 {
				fmt.Fprintln(out, "YES")
				continue
			}
			flag := false
			temp := x
			for i := 29; i >= 0; i-- {
				canbe := temp / int(math.Pow(2, float64(i)))
				if canbe > 0 {
					mini := Min(canbe, bits[i])
					temp -= mini * int(math.Pow(2, float64(i)))
					if temp == 0 {
						flag = true
						break
					}
				}
			}
			if flag {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
