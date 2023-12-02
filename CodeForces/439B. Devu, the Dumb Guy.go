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
    
    var n, x int
    fmt.Fscan(in, &n, &x)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &arr[i])
    }
    
    sort.Ints(arr)
    ans := 0
    for i := 0; i < n; i++ {
        ans += arr[i] * x
        if x > 1 {
            x--
        }
    }
    fmt.Fprintln(out, ans)
}

func main() {
    solve(os.Stdin, os.Stdout)
}
