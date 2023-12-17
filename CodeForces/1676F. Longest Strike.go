package main
 
import (
    "fmt"
    "os"
    "io"
    "bufio"
    "sort"
)
 
func Solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var t int
    fmt.Fscan(in, &t)
    for ; t > 0; t-- {
        var n, k int
        fmt.Fscan(in, &n, &k)
        arr := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Fscan(in, &arr[i])
        }
        
        mp := make(map[int]int)
        for _, v := range arr {
            mp[v]++
        }
        flag := false
        keys := []int{}
        kl := 0
        for ke, v := range mp {
            keys = append(keys, ke)
            kl++
            if v >= k && !flag {
                flag = true
            }
        }
        if !flag {
            fmt.Fprintln(out, -1)
            continue
        }
        
        sort.Ints(keys)
        l, r := 0, 1
        curr := 0
        ans := []int{keys[0], keys[0]}
    
        for l < kl && r < kl {
            if mp[keys[l]] < k {
                l++
                r = l + 1
            } else {
                if keys[r]-keys[r-1] == 1 && mp[keys[r]] >= k {
                    r++
                } else {
                    if curr < r-l && (mp[keys[l]] >= k && mp[keys[r-1]] >= k) {
                        curr = r - l
                        ans[0], ans[1] = keys[l], keys[r-1]
                    }
                    l = r
                    r++
                }
            }
        }
        if curr < r-l && (mp[keys[l]] >= k && mp[keys[r-1]] >= k) {
            curr = r - l
            ans[0], ans[1] = keys[l], keys[r-1]
        }
        fmt.Fprintln(out, ans[0], ans[1])
    }
}
 
func main() {
    Solve(os.Stdin, os.Stdout)
}
