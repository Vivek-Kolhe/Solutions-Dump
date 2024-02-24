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

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	if x > y {
		fmt.Fprintln(out, n)
		return
	}
	sort.Ints(arr)
	low, high := 0, n - 1
	for low <= high {
		mid := (low + high) >> 1
		if x >= arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low < 1 {
		fmt.Fprintln(out, 0)
		return
	}
	fmt.Fprintln(out, 1+(low-1)/2) 
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
