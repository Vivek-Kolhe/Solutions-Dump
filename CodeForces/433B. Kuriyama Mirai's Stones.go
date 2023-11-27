package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
    "sort"
)

func solve(_r io.Reader, _w io.Writer){
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var n, m int
    fmt.Fscan(in, &n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &arr[i])
    }
    fmt.Fscan(in, &m)
    
    pref := []int{0}
    for i := 0; i < n; i++ {
        pref = append(pref, arr[i]+pref[i])
    }
    sort.Ints(arr)
    prefs := []int{0}
    for i := 0; i < n; i++ {
        prefs = append(prefs, arr[i]+prefs[i])
    }
    
    for i := 0; i < m; i++ {
        var t, l, r int
        fmt.Fscan(in, &t, &l, &r)
        if t == 1 {
            fmt.Fprintln(out, pref[r]-pref[l-1])
        } else {
            fmt.Fprintln(out, prefs[r]-prefs[l-1])
        }
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
