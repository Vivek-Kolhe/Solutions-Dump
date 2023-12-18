package main
 
import (
    "fmt"
    "os"
    "io"
    "bufio"
    "strconv"
    "strings"
)

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func Solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var n int
    fmt.Fscan(in, &n)
    var s string
    fmt.Fscan(in, &s)
    maps := make([]int, 9)
    for i := 0; i < 9; i++ {
        fmt.Fscan(in, &maps[i])
    }
    
    digs := make([]string, n)
    for i, char := range s {
        dig := string(char)
        digs[i] = dig
    }
    
    for i := 0; i < n; i++ {
        num, _ := strconv.Atoi(digs[i])
        if maps[num-1]-num > 0 {
            for i < n && maps[num-1]-num >= 0 {
                num, _ = strconv.Atoi(digs[i])
                digs[i] = strconv.Itoa(Max(num, maps[num-1]))
                i++
            }
            break
        }
    }
    fmt.Fprintln(out, strings.Join(digs, ""))
}
 
func main() {
    Solve(os.Stdin, os.Stdout)
}
