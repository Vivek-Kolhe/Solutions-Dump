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

	var n, a, b int
	fmt.Fscan(in, &n)
	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		arr[i] = Pair{
			First:  a,
			Second: b,
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].First == arr[j].First {
			return arr[i].Second < arr[j].Second
		}
		return arr[i].First < arr[j].First
	})
	for i := 1; i < n; i++ {
		if arr[i].Second < arr[i-1].Second {
			fmt.Fprintln(out, arr[n-1].First)
			return
		}
	}
	fmt.Fprintln(out, arr[n-1].Second)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
