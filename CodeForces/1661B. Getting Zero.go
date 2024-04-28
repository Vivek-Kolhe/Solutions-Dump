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

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	mod := 32768
	ans := make([]int, n)
	mp := make(map[int]int)
	for i, v := range arr {
		q := make(chan Pair, 32768*2+1)
		seen := make([]bool, 32768+1)

		if _, ok := mp[v]; ok {
			ans[i] = mp[v]
			continue
		}
		q <- Pair{0, v}
		for len(q) > 0 {
			top := <-q
			moves, num := top.First, top.Second
			seen[num] = true

			if num == 0 {
				mp[v] = moves
				ans[i] = moves
				break
			}

			newNum := (num + 1) % mod
			if !seen[newNum] {
				q <- Pair{moves+1, newNum}
			}

			newNum = (2 * num) % mod
			if !seen[newNum] {
				q <- Pair{moves+1, newNum}
			}
		}
	}
	for _, v := range ans {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
