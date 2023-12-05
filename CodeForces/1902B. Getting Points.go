package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var tc int
    fmt.Fscan(in, &tc)
    for ; tc > 0; tc-- {
        var n, p, l, t int
        fmt.Fscan(in, &n, &p, &l, &t)
        
        tasks := (n + 6) / 7
        low, high := 0, n
        for low <= high {
            mid := (low + high) / 2
            if mid*l+min(tasks,2*mid)*t >= p {
                high = mid - 1
            } else {
                low = mid + 1
            }
        }
        fmt.Fprintln(out, n-high-1)
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
