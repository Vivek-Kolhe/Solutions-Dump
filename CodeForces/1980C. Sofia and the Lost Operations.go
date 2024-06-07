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
	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(in, &n)
		a, b := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &b[i])
		}
		fmt.Fscan(in, &m)
		d := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &d[i])
		}

		mp := make(map[int]int)
		for _, v := range d {
			mp[v]++
		}

		var flag, f bool
		for i := 0; i < n; i++ {
			if d[m-1] == b[i] {
				flag = true
			}
			if a[i] != b[i] {
				if _, ok := mp[b[i]]; !ok {
					f = true
					fmt.Fprintln(out, "NO")
					break
				}
				mp[b[i]]--
				if mp[b[i]] == 0 {
					delete(mp, b[i])
				}
			}
		}
		if !f {
			if len(mp) > 0 && !flag {
				fmt.Fprintln(out, "NO")
				continue
			}
			fmt.Fprintln(out, "YES")
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
