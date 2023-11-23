package main

import (
    "bufio"
    "os"
    "io"
    "fmt"
    "sort"
)

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
    sort.Ints(arr)
    
    if arr[0] != 1 {
        fmt.Fprintln(out, 1)
    } else {
        res := -1
        for i := 0; i < n-1; i++ {
            if arr[i+1]-arr[i] > 1 {
                res = arr[i] + 1
                break
            }
        }
        if res == -1 {
            res = arr[n-1] + 1
        }
        fmt.Fprintln(out, res)
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
