package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func BuildNum(bit []int) int {
	var ans int
	for i := 0; i < 32; i++ {
		if bit[i] > 0 {
			ans += (1 << i)
		}
	}
	return ans
}

func Func(mid, n int, arr []int) bool {
	bit := make([]int, 32)
	for i := 0; i < mid; i++ {
		for j := 0; j < 32; j++ {
			if arr[i]&(1<<j) > 0 {
				bit[j]++
			}
		}
	}

	checker := BuildNum(bit)
	for i := mid; i < n; i++ {
		for j := 0; j < 32; j++ {
			if arr[i-mid]&(1<<j) > 0 {
				bit[j]--
			}
		}

		for j := 0; j < 32; j++ {
			if arr[i]&(1<<j) > 0 {
				bit[j]++
			}
		}

		or := BuildNum(bit)
		if or != checker {
			return false
		}
	}
	return true
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		ans := n
		low, high := 1, n
		for low <= high {
			mid := (low + high) >> 1
			if Func(mid, n, arr) {
				ans = mid
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
