package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Pair struct {
	First  int
	Second int
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](nums ...T) T {
	res := nums[0]
	for _, v := range nums {
		if res > v {
			res = v
		}
	}
	return res
}

func Max[T Number](nums ...T) T {
	res := nums[0]
	for _, v := range nums {
		if res < v {
			res = v
		}
	}
	return res
}

func Abs[T Number](num T) T {
	if num < 0 {
		return num * -1
	}
	return num
}

func GCD(a, b int) int {
	for a > 0 && b > 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	if a == 0 {
		return b
	}
	return a
}

var q = make(chan Pair, 1000000)
var delRows = []int{0, 1, -1, 0}
var delCols = []int{-1, 0, 0, 1}

func bfs(n, m int, visited [][]bool, grid [][]int) int {
	var c int
	for len(q) > 0 {
		top := <-q
		c += grid[top.First][top.Second]
		for i := 0; i < 4; i++ {
			nrow := top.First + delRows[i]
			ncol := top.Second + delCols[i]
			if nrow >= 0 && nrow < n && ncol >= 0 && ncol < m && grid[nrow][ncol] != 0 && !visited[nrow][ncol] {
				q <- Pair{
					First:  nrow,
					Second: ncol,
				}
				visited[nrow][ncol] = true
			}
		}
	}
	return c
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(in, &n, &m)
		grid := make([][]int, n)
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			grid[i] = make([]int, m)
			visited[i] = make([]bool, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &grid[i][j])
			}
		}
		var ans int
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] != 0 && !visited[i][j] {
					q <- Pair{
						First: i,
						Second: j,
					}
					visited[i][j] = true
					ans = Max(ans, bfs(n, m, visited, grid))
				}
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
