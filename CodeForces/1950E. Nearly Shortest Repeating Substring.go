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

func ReverseString(str string) string {
	var reversed strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		reversed.WriteByte(str[i])
	}
	return reversed.String()
}

func RepeatString(str string, l int) string {
	var repeated strings.Builder
	for i := 0; i < l; i++ {
		repeated.WriteString(str)
	}
	return repeated.String()
}


func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)

		factors := make([]int, 0)
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				factors = append(factors, i)
			}
		}

		reverse := ReverseString(s)

		ans := n
		for i := 0; i < len(factors)-1; i++ {
			v := factors[i]
			var pc, sc int
			pref := RepeatString(s[:v], n/v)
			suff := RepeatString(reverse[:v], n/v)
			for j := 0; j < n; j++ {
				if pref[j] != s[j] {
					pc++
				}
				if suff[j] != reverse[j] {
					sc++
				}
			}
			if sc <= 1 || pc <= 1 {
				ans = v
				break
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
