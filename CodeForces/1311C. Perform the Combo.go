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

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(in, &n, &m)
		var s string
		fmt.Fscan(in, &s)
		arr := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &arr[i])
		}
		
		sort.Ints(arr)
		c := make([]int, n)
		res := make([]int, 26)
		for i, v := range arr {
			if i > 0 && arr[i] == arr[i-1] {
				continue
			}
			for j := v-1; j >= 0; j-- {
				if c[j] != 0 {
					break
				}
				c[j] = m - i
			}
		}
		for i := 0; i < n; i++ {
			res[s[i]-'a']++
			res[s[i]-'a'] += c[i]
		}
		for _, v := range res {
			fmt.Fprint(out, v, " ")
		}
		fmt.Fprintln(out)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
