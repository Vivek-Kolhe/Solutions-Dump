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

func BS(arr []int, key int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, pos, l, r int
	fmt.Fscan(in, &n, &pos, &l, &r)

	if l == 1 && r == n {
		fmt.Fprintln(out, 0)
		return
	}
	if (pos == r && l == 1) || (pos == l && r == n) {
		fmt.Fprintln(out, 1)
		return
	}
	if pos > r {
		if l == 1 {
			fmt.Fprintln(out, pos-r+1)
			return
		}
		fmt.Fprintln(out, pos-r+r-l+2)
	} else if pos < l {
		if r == n {
			fmt.Fprintln(out, l-pos+1)
			return
		}
		fmt.Fprintln(out, l-pos+r-l+2)
	} else {
		if l != 1 && r != n {
			fmt.Fprintln(out, Min(pos-l, r-pos)+r-l+2)
			return
		}
		if l == 1 {
			fmt.Fprintln(out, r-pos+1)
			return
		}
		if r == n {
			fmt.Fprintln(out, pos-l+1)
			return
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
