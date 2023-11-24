package main

import (
    "bufio"
    "os"
    "io"
    "fmt"
    "math"
)

func query(l, r int, arr []int) int {
    if l > 0 {
        return arr[r] - arr[l-1]
    }
    return arr[r]
}

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var t int
    fmt.Fscan(in, &t)
    for ; t > 0; t-- {
        var n, s int
        fmt.Fscan(in, &n, &s)
        arr := make([]int, n)
        ones := 0
        for i := 0; i < n; i++ {
            fmt.Fscan(in, &arr[i])
            if arr[i] == 1 {
                ones++
            }
        }
        
        if ones < s {
            fmt.Fprintln(out, -1)
            continue
        }
        pref := make([]int, n)
        pref[0] = arr[0]
        for i := 1; i < n; i++ {
            pref[i] = pref[i-1] + arr[i]
        }
        
        ans := math.MaxInt64
        for i := 0; i < n; i++ {
            low, high, pos := 0, n - 1, -1
            for low <= high {
                mid := (low + high) / 2
                if query(i, mid, pref) <= s {
                    pos = mid
                    low = mid + 1
                } else {
                    high = mid - 1
                }
            }
            if pos != -1 || query(i, pos, pref) == s {
                ans = int(math.Min(float64(ans), float64(n - (pos - i + 1))))
            }
        }
        fmt.Fprintln(out, ans)
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
