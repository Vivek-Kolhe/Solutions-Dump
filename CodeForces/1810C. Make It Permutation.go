package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, c, d int
		fmt.Fscan(in, &n, &c, &d)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		var ans int
		seen := make(map[int]int)
		for i := 0; i < n; i++ {
			if _, ok := seen[arr[i]]; ok {
				ans += c
				continue
			}
			seen[arr[i]] = 1
		}
		uniq := make([]int, 0)
		for k := range seen {
			uniq = append(uniq, k)
		}
		sort.Ints(uniq)
		m := len(uniq)
		curr := math.MaxInt64
		for i := m-1; i >= 0; i-- {
			missed := uniq[i] - i - 1
			curr = Min(curr, missed*d+(m-i-1)*c)
		}
		curr = Min(curr, m*c+d)
		fmt.Fprintln(out, ans+curr)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
