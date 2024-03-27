package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	idx := make(map[int][]int)
	for i := 0; i < n; i++ {
		idx[arr[i]] = append(idx[arr[i]], i)
	}
	ans := make([]Pair, 0)
	for k, v := range idx {
		if len(v) == 1 {
			ans = append(ans, Pair{
				First: k,
				Second: 0,
			})
			continue
		}
		var flag bool
		diff := v[1] - v[0]
		for i := 1; i < len(v); i++ {
			if v[i]-v[i-1] != diff {
				flag = true
				break
			}
		}
		if !flag {
			ans = append(ans, Pair{
				First: k,
				Second: diff,
			})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].First < ans[j].First
	})
	fmt.Fprintln(out, len(ans))
	for i := range ans {
		fmt.Fprintln(out, ans[i].First, ans[i].Second)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
