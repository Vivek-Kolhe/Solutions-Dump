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
 
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
 
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)
	var cnt int
	for i := 0; i < n; i++ {
		if arr[i] != sorted[i] {
			cnt++
		}
	}
	if cnt > 2 {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
	}
}
 
func main() {
	Solve(os.Stdin, os.Stdout)
}
