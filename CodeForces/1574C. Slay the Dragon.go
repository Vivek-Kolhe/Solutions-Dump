package main

import (
    "fmt"
    "sort"
    "os"
    "io"
    "bufio"
)

func Max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func Solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()

    var n, m, sum int
    fmt.Fscan(in, &n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &arr[i])
        sum += arr[i]
    }
    sort.Ints(arr)
    fmt.Fscan(in, &m)
    for ; m > 0; m-- {
        var x, y int
        fmt.Fscan(in, &x, &y)

        if x <= arr[0] {
            fmt.Fprintln(out, Max(0, x-arr[0])+Max(0, y-(sum-arr[0])))
        } else if x >= arr[n-1] {
            fmt.Fprintln(out, Max(0, x-arr[n-1])+Max(0, y-(sum-arr[n-1])))
        } else {
            low, high := 0, n - 1
            for low <= high {
                mid := (low + high) / 2
                if arr[mid] >= x {
                    high = mid - 1
                } else {
                    low = mid + 1
                }
            }
            small, big := arr[high], arr[high+1]
            csmall := Max(0, x-small) + Max(0, y-(sum-small))
            cbig := Max(0, x-big) + Max(0, y-(sum-big))
            fmt.Fprintln(out, Min(cbig, csmall))
        }
    }
}

func main() {
    Solve(os.Stdin, os.Stdout)
}
