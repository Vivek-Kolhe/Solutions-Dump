package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
    "sort"
)

func mapKeys(m map[int]bool) []int {
    keys := []int{}
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    
    var n int
    fmt.Fscan(in, &n)
    arr := make([]int, n)
    seen := make(map[int]bool)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &arr[i])
        seen[arr[i]] = true
    }
    
    keys := mapKeys(seen)
    if len(keys) > 3 {
        fmt.Fprintln(out, "NO")
    } else if len(keys) < 3 {
        fmt.Fprintln(out, "YES")
    } else {
        sort.Ints(keys)
        if keys[1]-keys[0] != keys[2]-keys[1] {
            fmt.Fprintln(out, "NO")
        } else {
            fmt.Fprintln(out, "YES")
        }
    }
}

func main() {
    solve(os.Stdin, os.Stdout)
}
