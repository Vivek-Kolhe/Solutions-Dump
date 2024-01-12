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

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, x, y, d int
		fmt.Fscan(in, &n, &x, &y, &d)

		if Abs(x-y)%d == 0 {
			fmt.Fprintln(out, Abs(x-y)/d)
			continue
		}
		if d > n {
			if y != n && y != 1 {
				fmt.Fprintln(out, -1)
				continue
			} 
		}
		var ans int
		var start, end bool
		if (y-1)%d == 0 {
			start = true
		}
		if (n-y)%d == 0 {
			end = true
		}
		if start && end {
			nl := int(math.Ceil(float64(Abs(x-1))/float64(d)))
			nr := int(math.Ceil(float64(Abs(x-n))/float64(d)))
			nl += ((y - 1) / d)
			nr += ((n - y) / d)
			fmt.Fprintln(out, Min(nl, nr))
			continue
		}
		if start {
			nl := int(math.Ceil(float64(Abs(x-1))/float64(d)))
			ans += (nl + (y - 1) / d)
			fmt.Fprintln(out, ans)
			continue
		}
		if end {
			nr := int(math.Ceil(float64(Abs(x-n))/float64(d)))
			ans += (nr + (n - y) / d)
			fmt.Fprintln(out, ans)
			continue
		}
		fmt.Fprintln(out, -1)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
