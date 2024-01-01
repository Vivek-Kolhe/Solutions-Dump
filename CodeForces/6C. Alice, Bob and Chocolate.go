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

	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	pref, suff := make([]int, n), make([]int, n)
	pref[0], suff[n-1] = arr[0], arr[n-1]
	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] + arr[i]
	}
	for i := n-2; i >= 0; i-- {
		suff[i] = suff[i+1] + arr[i]
	}
	var alice, bob int
	l, r := 0, n - 1
	for l <= r {
		if pref[l] < suff[r] {
			l++
			alice++
			continue
		}
		if pref[l] > suff[r] {
			r--
			bob++
			continue
		}
		if pref[l] == suff[r] {
			if l == r {
				alice++
			} else {
				r--
				alice++
				bob++
			}
			l++
		}
	}
	fmt.Fprintln(out, alice, bob)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
