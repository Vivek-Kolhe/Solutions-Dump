package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
)

func abs(a int) int {
    if a < 0 {
        return a * -1
    }
    return a
}

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var x1, y1, x2, y2 int
    fmt.Fscan(in, &x1, &y1, &x2, &y2)
    
    if x1 == x2 {
        dist := abs(y1 - y2)
        fmt.Fprintln(out, x1+dist, y1, x2+dist, y2)
    } else if y1 == y2 {
        dist := abs(x1 - x2)
        fmt.Fprintln(out, x1, y1+dist, x2, y2+dist)
    } else {
        hor := abs(x1 - x2)
        ver := abs(y1 - y2)
        if hor != ver {
            fmt.Fprintln(out, -1)
        } else {
            if y2 > y1 {
                fmt.Fprintln(out, x2, y1, x1, y2)
            } else {
                fmt.Fprintln(out, x1, y2, x2, y1)
            }
        }
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
