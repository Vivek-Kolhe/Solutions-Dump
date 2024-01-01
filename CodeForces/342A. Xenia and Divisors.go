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

	cnt := make([]int, 8)
	for _, v := range arr {
		cnt[v]++
	}
	if cnt[5] > 0 || cnt[7] > 0 {
		fmt.Fprintln(out, -1)
		return
	}
	if !(cnt[1] == cnt[2]+cnt[3] && cnt[2]+cnt[3] == cnt[4]+cnt[6]) {
		fmt.Fprintln(out, -1)
		return
	}
	var c int
	res := make([][]int, 0)
	for i := 0; i < cnt[4]; i++ {
		if cnt[2] < 1 {
			break
		}
		res = append(res, []int{1, 2, 4})
		c += 3
		cnt[2]--
	}
	for i := 0; i < cnt[6]; i++ {
		if cnt[3] < 1 {
			break
		}
		res = append(res, []int{1, 3, 6})
		c += 3
		cnt[3]--
	}
	for i := 0; i < cnt[6]; i++ {
		if cnt[2] < 1 {
			break
		}
		res = append(res, []int{1, 2, 6})
		c += 3
		cnt[2]--
	}
	if n != c {
		fmt.Fprintln(out, -1)
		return
	}
	for _, v := range res {
		for _, ele := range v {
			fmt.Fprint(out, ele, " ")
		}
		fmt.Fprintln(out)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
