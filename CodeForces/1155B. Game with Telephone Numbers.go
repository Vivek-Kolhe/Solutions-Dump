package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var n int
    var s string
    fmt.Fscan(in, &n)
    fmt.Fscan(in, &s)
    
    var cnt int
    for i := 0; i < n-10; i++ {
        if s[i] == '8' {
            cnt++
        }
    }
    if cnt > (n-10)/2 {
        fmt.Fprintln(out, "YES")
        return
    }
    fmt.Fprintln(out, "NO")
}

func main() {
    solve(os.Stdin, os.Stdout)
}
