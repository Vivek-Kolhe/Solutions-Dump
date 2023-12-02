package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
    "sort"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var t int
    fmt.Fscan(in, &t)
    for ; t > 0; t-- {
        var n, k int
        fmt.Fscan(in, &n, &k)
        s := make([]int, k)
        for i := 0; i < k; i++ {
            fmt.Fscan(in, &s[i])
        }
        arr := make([]int, k)
        if s[0] <= 0 {
            arr[0] = s[0] / (n - k + 1)
        } else {
            arr[0] = s[0] / (n - k + 1)
            if s[0]%(n-k+1) > 0 {
                arr[0]++
            }
        }
        
        for i := 1; i < k; i++ {
            arr[i] = s[i] - s[i-1]
        }
        
        if sort.IntsAreSorted(arr) {
            fmt.Fprintln(out, "Yes")
        } else {
            fmt.Fprintln(out, "No")
        }
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
