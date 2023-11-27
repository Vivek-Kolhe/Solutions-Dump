package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
)

func solve(_r io.Reader, _w io.Writer){
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var t int
    fmt.Fscan(in, &t)
    for ; t > 0; t-- {
        var n int
        fmt.Fscan(in, &n)
        var s string
        fmt.Fscan(in, &s)
        
        cnt := 0
        for _, v := range s {
            if v == '.' {
                cnt++
            }
        }
        
        flag := false
        for i := 1; i < n-1; i++ {
            if s[i-1] == '.' && s[i] == '.' && s[i+1] == '.' {
                fmt.Fprintln(out, 2)
                flag = true
                break
            }
        }
        
        if flag {
            continue
        }
        fmt.Fprintln(out, cnt)
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
