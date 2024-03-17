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

func DFS(n, row, col int, grid []string, visited [][]bool) {
	delRows := []int{0, 1, -1, 0}
    delCols := []int{-1, 0, 0, 1}
	visited[row][col] = true

	if row%2 != col%2 {
		if grid[row][col] == '<' {
			DFS(n, row, col-1, grid, visited)
		} else {
			DFS(n, row, col+1, grid, visited)
		}
		return
	}
	for i := 0; i < 4; i++ {
		nrow := row + delRows[i]
		ncol := col + delCols[i]
		if nrow >= 0 && nrow <= 1 && ncol >= 0 && ncol < n && !visited[nrow][ncol] {
			DFS(n, nrow, ncol, grid, visited)
		}
	}
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
		grid := make([]string, 2)
		var row string
		for i := 0; i < 2; i++ {
			fmt.Fscan(in, &row)
			grid[i] = row
		}

		visited := make([][]bool, 2)
		for i := 0; i < 2; i++ {
			visited[i] = make([]bool, n)
		}
		DFS(n, 0, 0, grid, visited)
		if visited[1][n-1] {
			fmt.Fprintln(out, "YES")
			continue
		}
		fmt.Fprintln(out, "NO")
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
