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

func DFS(node, c int, visited []bool, adj [][]int, mp map[int]int) {
	visited[node] = true
	mp[node] = c
	for _, n := range adj[node] {
		if !visited[n] {
			DFS(n, c, visited, adj, mp)
		}
	}
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	adj := make([][]int, n+1)
	var l int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &l)
		temp := make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Fscan(in, &temp[j])
		}
		for j := 0; j < l-1; j++ {
			adj[temp[j]] = append(adj[temp[j]], temp[j+1])
			adj[temp[j+1]] = append(adj[temp[j+1]], temp[j])
		}
	}
	
	var c int
	visited := make([]bool, n+1)
	mp := make(map[int]int)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			c++
			DFS(i, c, visited, adj, mp)
		}
	}
	
	lengths := make(map[int]int)
	for _, v := range mp {
		lengths[v]++
	}
	
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, lengths[mp[i]], " ")
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
