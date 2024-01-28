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

	var n, q, sum, prev int
	fmt.Fscan(in, &n, &q)
	arr := make([]int, n)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		sum += arr[i]
		mp[i] = arr[i]
	}
	for ; q > 0; q-- {
		var t, pos, new int
		fmt.Fscan(in, &t)
		if t == 1 {
			fmt.Fscan(in, &pos, &new)
			pos--
			if _, ok := mp[pos]; !ok {
				mp[pos] = prev
			}
			sum = sum - mp[pos] + new
			mp[pos] = new
			fmt.Fprintln(out, sum)
		}
		if t == 2 {
			fmt.Fscan(in, &new)
			mp = make(map[int]int)
			sum = new * n
			prev = new
			fmt.Fprintln(out, sum)
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
