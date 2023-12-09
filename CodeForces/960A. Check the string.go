package main

import (
    "fmt"
    "io"
    "os"
    "bufio"
)

func Group(s string, n int) []int {
    res := []int{}
    st := string(s[0])
    curr := 1
    for i := 1; i < n; i++ {
        if s[i] == st[curr-1] {
            curr++
            st += string(s[i])
        } else {
            res = append(res, curr)
            curr = 1
            st = string(s[i])
        }
    }
    res = append(res, curr)
    return res
}

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var s string
    fmt.Fscan(in, &s)
    
    n := len(s)
    if s[0] != 'a' || s[n-1] != 'c' {
        fmt.Fprintln(out, "NO")
        return
    }
    cnt := Group(s, n)
    if len(cnt) != 3 {
        fmt.Fprintln(out, "NO")
        return
    }
    if cnt[0] != cnt[2] && cnt[1] != cnt[2] {
        fmt.Fprintln(out, "NO")
        return
    }
    fmt.Fprintln(out, "YES")
}

func main() {
    solve(os.Stdin, os.Stdout)
}
