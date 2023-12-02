package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
    "sort"
)

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var n int
    fmt.Fscan(in, &n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &arr[i])
    }
    
    ans := 1
    sort.Ints(arr)
    l, r := 0, 0
    for l < n && r < n {
        if arr[r]-arr[l] <= 5 {
            ans = max(ans, r-l+1)
            r++
        } else {
            l++
        }
    }
    
    fmt.Fprintln(out, ans)
}

func main() {
    solve(os.Stdin, os.Stdout)
}
