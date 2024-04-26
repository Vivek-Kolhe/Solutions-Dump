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
		var n int
		fmt.Fscan(in, &n)
		grid := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &grid[i])
		}

		var flag bool
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == '1' {
					var nrow, ncol int
					ncol = j + 1
					nrow = i + 1
					if ncol >= 0 && ncol <= n {
						if ncol == n || grid[i][ncol] == '1' {
							continue
						} 
					}
					if nrow >= 0 && nrow <= n {
						if nrow == n || grid[nrow][j] == '1' {
							continue
						}
					}
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
		if flag {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
