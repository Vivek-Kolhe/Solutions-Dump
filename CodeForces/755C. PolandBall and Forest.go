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

func DFS(node int, visited []bool, adj [][]int) {
	visited[node] = true
	for _, v := range adj[node] {
		if !visited[v] {
			DFS(v, visited, adj)
		}
	}
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, edge int
	fmt.Fscan(in, &n)
	adj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &edge)
		adj[i] = append(adj[i], edge)
		adj[edge] = append(adj[edge], i)
	}
	
	var ans int
	visited := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			ans++
			DFS(i, visited, adj)
		}
	}
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
