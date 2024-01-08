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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &arr[i])
	}
	sort.Ints(arr)
	var curr, ans int
	var flag bool
	curr = 1
	for ; k > 0; k-- {
		var a, b int
		fmt.Fscan(in, &a, &b)
		
		if flag {
			continue
		}
		ind := BS(arr, curr)
		if ind != -1 {
			ans = arr[ind]
			flag = true
		}
		if curr == a {
			curr = b
		} else if curr == b {
			curr = a
		}
	}
	if ans == 0 {
		fmt.Fprintln(out, curr)
	} else {
		fmt.Fprintln(out, ans)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
