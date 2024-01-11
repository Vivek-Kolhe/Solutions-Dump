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
	low, high := 0, len(arr)-1
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

	lengths := make([]Pair, 4)
	choices := []string{"A", "B", "C", "D"}
	var s string
	for i := 0; i < 4; i++ {
		fmt.Fscan(in, &s)
		lengths[i] = Pair{
			First: len(s[2:]),
			Second: i,
		}
	}

	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i].First < lengths[j].First
	})
	var good, c, g int
	for i := 1; i < 4; i++ {
		if lengths[0].First <= lengths[i].First/2 {
			c++
		}
	}
	if c == 3 {
		good++
		g = lengths[0].Second
	}
	c = 0
	for i := 0; i < 3; i++ {
		if lengths[3].First >= lengths[i].First*2 {
			c++
		}
	}
	if c == 3 {
		good++
		g = lengths[3].Second
	}
	if good != 1 {
		fmt.Fprintln(out, "C")
	} else {
		fmt.Fprintln(out, choices[g])
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
