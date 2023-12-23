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

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		pwr := make([]Pair, n)
		for i := 0; i < n; i++ {
			var l int
			fmt.Fscan(in, &l)
			cave := make([]int, l)
			currp := 0
			for j := 0; j < l; j++ {
				fmt.Fscan(in, &cave[j])
				currp = Max(currp, cave[j]-j+1)
			}
			pwr[i] = Pair{
				First:  currp,
				Second: l,
			}
		}
		sort.Slice(pwr, func(i, j int) bool {
			return pwr[i].First < pwr[j].First
		})
		low, high := pwr[0].First, pwr[n-1].First
		for low <= high {
			mid := (low + high) >> 1
			curr := mid
			flag := true
			for i := 0; i < n; i++ {
				if curr >= pwr[i].First {
					curr += pwr[i].Second
				} else {
					flag = false
					break
				}
			}
			if flag {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		fmt.Fprintln(out, low)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
