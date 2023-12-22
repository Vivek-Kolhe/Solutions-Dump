package main
 
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
 
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
 
	var n, m int
	fmt.Fscan(in, &n, &m)
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &arr[i])
	}
 
	ans := make([]int, 0)
	mp := make(map[int][]int)
	for i, v := range arr {
		mp[v] = append(mp[v], i)
	}
Loop:
	for {
		maxi := 0
		s := 0
		for k, v := range mp {
			if len(mp[k]) == 0 {
				break Loop
			}
			if v[0] > maxi {
				maxi = v[0]
			}
			if len(mp[k]) > 0 {
				mp[k] = v[1:]
			} else {
				mp[k] = []int{}
			}
			s += k
		}
		if s == (n*(n+1))/2 {
			ans = append(ans, maxi)
		}
	}
	itr := 0
	for i := 0; i < m; i++ {
		if itr < len(ans) && i == ans[itr] {
			fmt.Fprint(out, "1")
			itr++
		} else {
			fmt.Fprint(out, "0")
		}
	}
	fmt.Fprintln(out)
}
 
func main() {
	Solve(os.Stdin, os.Stdout)
}
