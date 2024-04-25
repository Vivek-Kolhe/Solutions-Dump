package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
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
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		closest := make([]int, n)
		closest[0] = 1
		for i := 1; i < n-1; i++ {
			p := Abs(arr[i]-arr[i-1])
			n := Abs(arr[i+1]-arr[i])
			if p < n {
				closest[i] = i - 1
			} else {
				closest[i] = i + 1
			}
		}
		closest[n-1] = n - 2

		pref := make([]int, n+1)
		for i := 1; i < n; i++ {
			if closest[i-1] == i {
				pref[i+1] = pref[i] + 1
			} else {
				pref[i+1] = pref[i] + Abs(arr[i]-arr[i-1])
			}
		}
		
		suff := make([]int, n+1)
		for i := n-2; i >=0 ; i-- {
			if closest[i+1] == i {
				suff[i] = suff[i+1] + 1
			} else {
				suff[i] = suff[i+1] + Abs(arr[i]-arr[i+1])
			}
		}
		
		var m int
		fmt.Fscan(in, &m)
		for ; m > 0; m-- {
			var l, r int
			fmt.Fscan(in, &l, &r)
			if l == r {
				fmt.Fprintln(out, 0)
			} else if l > r {
				r--; l--
				fmt.Fprintln(out, suff[r]-suff[l])
			} else {
				fmt.Fprintln(out, pref[r]-pref[l])
			}
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
