package main

import (
    "bufio"
    "os"
    "fmt"
    "io"
    "sort"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var n, k, x int
    fmt.Fscan(in, &n, &k, &x)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &arr[i])
    }
    
    sort.Ints(arr)
    diffs := []int{}
    for i := 0; i < n-1; i++ {
        diff := arr[i+1] - arr[i]
        if diff > x {
            diffs = append(diffs, diff)
        }
    }
    sort.Ints(diffs)
    cnt := 1
    for i := 0; i < len(diffs); i++ {
        need := (diffs[i] - 1) / x
        if need > k {
            cnt++
        } else {
            k -= need
        }
    }
    fmt.Fprintln(out, cnt)
}

func main() {
    solve(os.Stdin, os.Stdout)
}
