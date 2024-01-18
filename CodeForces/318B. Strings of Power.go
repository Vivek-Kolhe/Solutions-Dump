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

	var s string
	fmt.Fscan(in, &s)

	heavy, metal := make([]int, 0), make([]int, 0)
	for i := 0; i < len(s)-5+1; i++ {
		if s[i:i+5] == "heavy" {
			heavy = append(heavy, i)
		}
		if s[i:i+5] == "metal" {
			metal = append(metal, i)
		}
	}
	var ans int
	for _, v := range metal {
		low, high := 0, len(heavy) - 1
		for low <= high {
			mid := (low + high) >> 1
			if heavy[mid] <= v {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		ans += low
	}
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
