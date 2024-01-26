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

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var s string
	fmt.Fscan(in, &s)

	if len(s) < 26 {
		fmt.Fprintln(out, -1)
		return
	}
	var flag bool
	var start, end int
	seen := make(map[byte]int)
	var curr int
	for i := 0; i < 26; i++ {
		if s[i] != '?' {
			seen[s[i]]++
		} else {
			curr++
		}
	}
	if len(seen) == 26 || curr+len(seen) == 26 {
		for i := 0; i < 26; i++ {
			if s[i] != '?' {
				fmt.Fprint(out, string(s[i]))
			} else {
				for _, ch := range letters {
					if _, ok := seen[byte(ch)]; !ok {
						fmt.Fprint(out, string(ch))
						seen[byte(ch)]++
						break
					}
				}
			}
		}
		for i := 26; i < len(s); i++ {
			if s[i] == '?' {
				fmt.Fprint(out, "A")
				continue
			}
			fmt.Fprint(out, string(s[i]))
		}
		fmt.Fprintln(out)
		return
	}
	for i := 1; i < len(s)-26+1; i++ {
		if s[i-1] == '?' {
			curr--
		}
		if s[25+i] == '?' {
			curr++
		}
		if s[i-1] != '?' {
			seen[s[i-1]]--
			if seen[s[i-1]] <= 0 {
				delete(seen, s[i-1])
			}
		}
		if s[25+i] != '?' {
			seen[s[25+i]]++
		}
		if (len(seen) == 26 && curr == 0) || curr+len(seen) == 26 {
			start, end = i, 25 + i
			flag = true
			break
		}
	}
	if !flag {
		fmt.Fprintln(out, -1)
		return
	} else {
		for i := 0; i < start; i++ {
			if s[i] == '?' {
				fmt.Fprint(out, "A")
				continue
			}
			fmt.Fprint(out, string(s[i]))
		}
		for i := start; i <= end; i++ {
			if s[i] != '?' {
				fmt.Fprint(out, string(s[i]))
			} else {
				for _, ch := range letters {
					if _, ok := seen[byte(ch)]; !ok {
						fmt.Fprint(out, string(ch))
						seen[byte(ch)]++
						break
					}
				}
			}
		}
		for i := end+1; i < len(s); i++ {
			if s[i] == '?' {
				fmt.Fprint(out, "A")
				continue
			}
			fmt.Fprint(out, string(s[i]))
		}
		fmt.Fprintln(out)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
