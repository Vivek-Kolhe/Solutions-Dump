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

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(in, &n, &k)
		var s string
		fmt.Fscan(in, &s)

		oneidx := make([]int, 0)
		for i, v := range s {
			if v == '1' {
				oneidx = append(oneidx, i)
			}
		}
		var ans int
		if len(oneidx) < 1 {
			for i := 0; i < n; i += (k+1) {
				ans++
			}
			fmt.Fprintln(out, ans)
			continue
		}
		if oneidx[0] != 0 {
			if oneidx[0] > k {
				ans++
			}
			oneidx = append([]int{0}, oneidx... )
		}
		if oneidx[len(oneidx)-1] != n-1 {
			if n-oneidx[len(oneidx)-1]-1 > k {
				ans++
			}
			oneidx = append(oneidx, n-1)
		}
		for i := 0; i < len(oneidx)-1; i++ {
			for j := oneidx[i]+k+1; j <= oneidx[i+1]-k-1; j += (k+1) {
				ans++
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
