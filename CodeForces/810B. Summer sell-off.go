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

	var a, b, n, f int
	fmt.Fscan(in, &n, &f)
	k, l := make([]int, n), make([]int, n)
	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		k[i], l[i] = a, b
		arr[i] = Pair{
			First: Min(2*a, b) - Min(a, b),
			Second: i,
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].First > arr[j].First
	})
	var ans int
	for i := 0; i < f; i++ {
		pos := arr[i].Second
		ans += Min(2*k[pos], l[pos])
	}
	for i := f; i < n; i++ {
		pos := arr[i].Second
		ans += Min(k[pos], l[pos])
	}
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
