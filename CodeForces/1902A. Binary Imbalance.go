package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var t int
    fmt.Fscan(in, &t)
    for ; t > 0; t-- {
        var n int
        var s string
        fmt.Fscan(in, &n)
        fmt.Fscan(in, &s)
        
        cnt := 0
        for i := 0; i < n; i++ {
            if s[i] == '0' {
                cnt++
            }
        }
        if cnt > n-cnt {
            fmt.Fprintln(out, "YES")
        } else {
            flag := false
            for i := 0; i < n-1; i++ {
                if (s[i] == '0' && s[i+1] == '1') || (s[i] == '1' && s[i+1] == '0') {
                    fmt.Fprintln(out, "YES")
                    flag = true
                    break
                }
            }
            if !flag {
                fmt.Fprintln(out, "NO")
            }
        }
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
