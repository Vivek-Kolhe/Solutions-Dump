package main

import (
    "bufio"
    "os"
    "fmt"
    "io"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var t int
    fmt.Fscan(in, &t)
    for ; t > 0; t-- {
        var n int
        fmt.Fscan(in, &n)
        arr := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Fscan(in, &arr[i])
        }
        
        mini := arr[0]
        ind := 0
        for i := 1; i < n; i++ {
            if mini > arr[i] {
                mini = arr[i]
                ind = i
            }
        }
        
        res := ind
        for i := ind; i < n-1; i++ {
            if arr[i] > arr[i+1] {
                res = -1
                break
            }
        }
        fmt.Fprintln(out, res)
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
