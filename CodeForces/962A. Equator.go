package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	var n, val int
	fmt.Fscan(in, &n)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &val)
		arr[i] = arr[i-1] + val
	}

	low, high := 1, n
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] >= int(math.Ceil(float64(arr[n])/2)) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Fprintln(out, low)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
