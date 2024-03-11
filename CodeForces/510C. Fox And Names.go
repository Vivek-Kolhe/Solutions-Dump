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

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	
	adj := make([][]byte, 26)
	for i := 0; i < n-1; i++ {
		var flag bool
		curr, next := arr[i], arr[i+1]
		for j := 0; j < Min(len(curr), len(next)); j++ {
			if curr[j] != next[j] {
				adj[curr[j]-'a'] = append(adj[curr[j]-'a'], next[j]-'a')
				flag = true
				break
			}
		}
		if !flag {
			if len(curr) > len(next) {
				fmt.Fprintln(out, "Impossible")
				return
			}
		}
	}
	q := make(chan int, 26)
	inDegree := make([]int, 26)
	topo := make([]int, 0)
	for i := 0; i < 26; i++ {
		for _, v := range adj[i] {
			inDegree[v]++
		}
	}
	for i := 0; i < 26; i++ {
		if inDegree[i] == 0 {
			q <- i
		}
	}
	for len(q) > 0 {
		top := <- q
		topo = append(topo, top)
		for _, v := range adj[top] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q <- int(v)
			}
		}
	}
	if len(topo) != 26 {
		fmt.Fprintln(out, "Impossible")
		return
	}
	for i := 0; i < len(topo); i++ {
		fmt.Fprint(out, string(topo[i]+'a'))
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
