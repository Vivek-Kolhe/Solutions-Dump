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

	var n, k int
	fmt.Fscan(in, &n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	sort.Ints(arr)
	low, high := n / 2, n - 1
	var fcost int
	for low <= high {
		mid := (low + high) >> 1
		var cost int
		for i := n/2; i < mid; i++ {
			cost += (arr[mid] - arr[i])
		}
		if cost <= k {
			fcost = cost
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	ans := arr[low-1] + ((k - fcost) / (low - n / 2)) 
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
