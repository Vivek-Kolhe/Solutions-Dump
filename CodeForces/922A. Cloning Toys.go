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
 
	var x, y int
	fmt.Fscan(in, &x, &y)
 
	if (y > x && Abs(x-y) > 1) || y == 0 {
		fmt.Fprintln(out, "NO")
		return
	}
	if y == 1 {
		if x > 0 {
			fmt.Fprintln(out, "NO")
			return
		}
		fmt.Fprintln(out, "YES")
		return
	}
	needy := y - 1
	needx := x - needy
	if needx%2 == 0 {
		fmt.Fprintln(out, "YES")
		return
	}
	fmt.Fprintln(out, "NO")
}
 
func main() {
	Solve(os.Stdin, os.Stdout)
}
