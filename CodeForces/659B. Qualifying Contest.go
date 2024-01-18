package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type Pair struct {
	First  string
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

	var n, m int
	fmt.Fscan(in, &n, &m)
	mp := make(map[int][]Pair)
	for i := 0; i < n; i++ {
		var region, score int
		var name string
		fmt.Fscan(in, &name, &region, &score)
		p := Pair{
			First: name,
			Second: score,
		}
		mp[region] = append(mp[region], p)
	}

	for k := 1; k <= m; k++ {
		sort.Slice(mp[k], func(i, j int) bool {
			return mp[k][i].Second > mp[k][j].Second
		})
	}
	ans := make(map[int][]string)
	for i := 1; i <= m; i++ {
		curr := mp[i]
		if len(curr) == 2 {
			ans[i] = append(ans[i], curr[0].First, curr[1].First)
			continue
		}
		ans[i] = append(ans[i], curr[0].First)
		if curr[1].Second != curr[2].Second {
			ans[i] = append(ans[i], curr[1].First)
		} else {
			ans[i] = []string{"?"}
		}
	}
	for i := 1; i <= m; i++ {
		curr := ans[i]
		if len(curr) < 2 {
			fmt.Fprintln(out, curr[0])
		} else {
			fmt.Fprintln(out, curr[0], curr[1])
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
