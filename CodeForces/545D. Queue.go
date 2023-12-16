package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "sort"
)

func Solve(_r io.Reader, _w io.Writer) {
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
    s, cnt := 0, 0
    for i := 0; i < n; i++ {
        if s > arr[i] {
            cnt++
        } else {
            s += arr[i]
        }
    }
    fmt.Fprintln(out, n-cnt)
}

func main() {
    Solve(os.Stdin, os.Stdout)
}
