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
	low, high := 0, len(arr) - 1
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

	var a, b, x int
	fmt.Fscan(in, &a, &b, &x)

	if a == 1 {
		if x == 1 {
			fmt.Fprint(out, "0")
		} else if x == 2 {
			fmt.Fprint(out, "10")
			b--
		}
		for i := 0; i < b; i++ {
			fmt.Fprint(out, "1")
		}
		fmt.Fprintln(out)
		return
	}
	if b == 1 {
		if x == 1 {
			fmt.Fprint(out, 1)
		} else if x == 2 {
			fmt.Fprint(out, "01")
			a--
		}
		for i := 0; i < a; i++ {
			fmt.Fprint(out, "0")
		}
		fmt.Fprintln(out)
		return
	}
	e, s := "1", "0"
	if a < b {
		s, e = e, s
		a, b = b, a
	}
	var last string
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			fmt.Fprint(out, s)
			last = s
			a--
		} else {
			fmt.Fprint(out, e)
			last = e
			b--
		}
	}
	if last == e {
		for i := 0; i < b; i++ {
			fmt.Fprint(out, e)
		}
		for i := 0; i < a; i++ {
			fmt.Fprint(out, s)
		}
	} else {
		for i := 0; i < a; i++ {
			fmt.Fprint(out, s)
		}
		for i := 0; i < b; i++ {
			fmt.Fprint(out, e)
		}
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
