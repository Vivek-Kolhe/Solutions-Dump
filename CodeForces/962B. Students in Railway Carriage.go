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

	var n, a, b int
	var s string
	fmt.Fscan(in, &n, &a, &b)
	fmt.Fscan(in, &s)

	seats := make([]int, 0)
	l, r := 0, 1
	for r < n {
		if s[l] == s[r] && s[l] == '.' {
			r++
		} else {
			if s[l] == '.' {
				seats = append(seats, r-l)
			}
			l = r
			r = l + 1
		}
	}
	if s[l] == '.' {
		seats = append(seats, r-l)
	}
	var ans int
	for _, seatLength := range seats {
		if seatLength%2 == 1 {
			if a > b {
				an := Min(a, (seatLength/2)+1)
				bn := Min(b, seatLength/2)
				a -= an
				b -= bn
				ans += (an + bn)
			} else {
				an := Min(a, seatLength/2)
				bn := Min(b, (seatLength/2)+1)
				a -= an
				b -= bn
				ans += (an + bn)
			}
		} else {
			an := Min(a, seatLength/2)
			bn := Min(b, seatLength/2)
			a -= an
			b -= bn
			ans += (an + bn)
		}
	}
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
