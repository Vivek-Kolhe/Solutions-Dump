package main
 
import (
    "fmt"
    "os"
    "io"
    "bufio"
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
    
    even, odd := 0, 0
    ei, oi := -1, -1
    for i := 0; i < n; i++ {
        if arr[i]%2 == 0 {
            even++
            ei = i + 1
        } else {
            odd++
            oi = i + 1
        }
    }
    if even > odd {
        fmt.Fprintln(out, oi)
    } else {
        fmt.Fprintln(out, ei)
    }
}
 
func main() {
    Solve(os.Stdin, os.Stdout)
}
