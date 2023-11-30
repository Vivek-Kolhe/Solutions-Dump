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
        var n, budget int
        fmt.Fscan(in, &n, &budget)
        arr := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Fscan(in, &arr[i])
        }
        
        sort.Ints(arr)
        for i := 1; i < n; i++ {
            arr[i] += arr[i-1]
        }
        ans, curr := 0, 0
        for i := n-1; i >= 0; i-- {
            arr[i] += curr * (i + 1)
            if budget-arr[i] >= 0 {
                days := ((budget - arr[i]) / (i + 1)) + 1
                ans += days * (i + 1)
                curr += days
            }
        }
        fmt.Fprintln(out, ans)
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
