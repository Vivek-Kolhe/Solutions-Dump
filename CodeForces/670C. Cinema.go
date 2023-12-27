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

	var n, m int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	fmt.Fscan(in, &m)
	audio, subs := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &audio[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &subs[i])
	}

	cnt := make(map[int]int)
	for _, v := range arr {
		cnt[v]++
	}
	ans, curra, currs := 0, audio[0], subs[0] 
	for i := 1; i < m; i++ {
		if cnt[audio[i]] > cnt[curra] {
			ans = i
			curra = audio[i]
			currs = subs[i]
		} else {
			if cnt[audio[i]] == cnt[curra] {
				if cnt[currs] < cnt[subs[i]] {
					ans = i
					curra = audio[i]
					currs = subs[i]
				}
			}
		}
	}
	fmt.Fprintln(out, ans+1)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
