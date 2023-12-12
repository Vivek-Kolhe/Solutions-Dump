package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
    "sort"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var n, k int
    fmt.Fscan(in, &n, &k)
    a := make([]int, n)
    b := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &a[i])
    }
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &b[i])
    }
    
    arr := make([][]int, n)
    for i := 0; i < n; i++ {
        arr[i] = []int{a[i], b[i]}
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i][0] - arr[i][1] < arr[j][0] - arr[j][1]
    })
    
    var ans, i int
    for i < n && (k > 0 || arr[i][0]-arr[i][1] < 1) {
        ans += arr[i][0]
        i++
        k--
    }
    for j := i; j < n; j++ {
        ans += arr[j][1]
    }
    fmt.Fprintln(out, ans)
}

func main() {
    solve(os.Stdin, os.Stdout)
}
