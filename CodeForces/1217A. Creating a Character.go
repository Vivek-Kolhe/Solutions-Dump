package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
)

func Solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var t int
    fmt.Fscan(in, &t)
    for ; t > 0; t-- {
        var str, intel, exp int
        fmt.Fscan(in, &str, &intel, &exp)
        
        if str+exp <= intel {
            fmt.Fprintln(out, 0)
            continue
        }
        low, high := 0, exp
        ans := low
        for low <= high {
            mid := (low + high) >> 1
            if str+mid > intel+exp-mid {
                ans = mid
                high = mid - 1
            } else {
                low = mid + 1
            }
        }
        fmt.Fprintln(out, exp-ans+1)
    }
}

func main() {
    Solve(os.Stdin, os.Stdout)
}
