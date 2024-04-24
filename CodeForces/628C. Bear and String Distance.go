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

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)

	var ans string
	for i := 0; i < n; i++ {
		back := s[i] - 'a'
		forward := 'z' - s[i]
		maxi := max(forward, back)

		if k >= int(maxi) {
			if back > forward {
				ans += "a"
			} else {
				ans += "z"
			}
			k -= int(maxi)
		} else {
			if back > forward {
				ans += string(s[i] - byte(k))
			} else {
				ans += string(s[i] + byte(k))
			}
			k = 0
		}
		if k == 0 {
			break
		}
	}
	if k > 0 {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprint(out, ans)
	for i := len(ans); i < n; i++ {
		fmt.Fprint(out, string(s[i]))
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
