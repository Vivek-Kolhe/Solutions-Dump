package main

import "fmt"

func main() {
  var n, m int
  fmt.Scanf("%d %d", &n, &m)
  if m % n != 0{
    fmt.Println(-1)
  } else {
    count := 0
    var d int = m / n
    for ; d % 2 == 0; {
      count ++
      d = d / 2
    }
    for ; d % 3 == 0; {
      count ++
      d = d / 3
    }
    if d != 1 {
      fmt.Println(-1)
    } else {
      fmt.Println(count)
    }
  }
}
