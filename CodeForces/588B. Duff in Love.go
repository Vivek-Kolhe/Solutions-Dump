package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	var n int
	fmt.Fscan(in, &n)

	divs := []int{}
	perfsq := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i == i {
				divs = append(divs, i)
			} else {
				divs = append(divs, n/i, i)
			}
		}
	}
	odd := 3
	for i := 1; i+odd <= n; {
		perfsq = append(perfsq, i+odd)
		i += odd
		odd += 2
	} 
	sort.Slice(divs, func(i, j int) bool {
		return divs[i] > divs[j]
	})
	for _, v := range divs {
		flag := true
		for _, sq := range perfsq {
			if v%sq == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Fprintln(out, v)
			break
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
