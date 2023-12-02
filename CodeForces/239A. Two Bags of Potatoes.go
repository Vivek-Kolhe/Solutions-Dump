package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var y, k, n int
    fmt.Fscan(in, &y, &k, &n)
    
    if n <= y {
        fmt.Fprintln(out, -1)
        return
    }
    possx := n - y
    var first int
    if y < k || y > k {
        first = k - (y % k)
        if possx >= first {
            fmt.Fprint(out, first, " ")
            first += k
        } else {
            fmt.Fprintln(out, -1)
            return
        }
    } else {
        fmt.Fprint(out, k, " ")
        first = k + k
    }
    if first == 0 {
        fmt.Fprintln(out, -1)
        return
    }
    for first <= possx {
        fmt.Fprint(out, first, " ")
        first += k
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
