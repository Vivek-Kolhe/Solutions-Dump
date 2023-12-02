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
    
    var c1, c2, c3, c4 int
    fmt.Fscan(in, &c1, &c2, &c3, &c4)
    var n, m int
    fmt.Fscan(in, &n, &m)
    bus, trolley := make([]int, n), make([]int, m)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &bus[i])
    }
    for i := 0; i < m; i++ {
        fmt.Fscan(in, &trolley[i])
    }
    
    sumb, sumt := 0, 0
    for i := 0; i < n; i++ {
        sumb += min(bus[i] * c1, c2)
    }
    for i := 0; i < m; i++ {
        sumt += min(trolley[i] * c1, c2)
    }
    sumb, sumt = min(sumb, c3), min(sumt, c3)
    
    ans := min(sumb+sumt, c4)
    fmt.Fprintln(out, ans)
}

func main() {
    solve(os.Stdin, os.Stdout)
}
