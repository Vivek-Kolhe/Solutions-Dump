package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		if n == 2 {
			if arr[0] > arr[1] {
				fmt.Fprintln(out, "NO")
			} else {
				fmt.Fprintln(out, "YES")
			}
			continue
		}

		stack := list.New()
		for i := 0; i < n; i++ {
			if stack.Len() == 0 {
				stack.PushBack(arr[i])
				continue
			} else {
				last := -1
				for stack.Len() > 0 && stack.Back().Value.(int) < arr[i] {
					last = stack.Back().Value.(int)
					stack.Remove(stack.Back())
				}
				if last != -1 {
					stack.PushBack(min(last, arr[i]))
				} else {
					stack.PushBack(arr[i])
				}
			}
		}
		if stack.Len() == 1 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
