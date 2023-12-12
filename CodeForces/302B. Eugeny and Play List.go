package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var n, m int
    fmt.Fscan(in, &n, &m)
    var t, c int
    fmt.Fscan(in, &t, &c)
    pref := []int{t*c}
    for i := 1; i < n; i++ {
        fmt.Fscan(in, &t, &c)
        pref = append(pref, pref[len(pref)-1]+(t*c))
    }
    for i := 0; i < m; i++ {
        var moment int
        fmt.Fscan(in, &moment)
        low, high := 0, n - 1
        for low <= high {
            mid := (low + high) / 2
            if pref[mid] < moment {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
        fmt.Fprintln(out, low+1)
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
