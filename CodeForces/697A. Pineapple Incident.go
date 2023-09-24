package main

import "fmt"

func main() {
  var t, s, x int
  fmt.Scanf("%d %d %d", &t, &s, &x)
  if (x == t) {
    fmt.Println("YES")
  } else if (x < t) {
    fmt.Println("NO")
  } else {
    x -= t
    if (x == 1) {
      fmt.Println("NO")
    } else if (x % s == 0 || (x - 1) % s == 0) {
      fmt.Println("YES")
    } else {
      fmt.Println("NO")
    }
  }
}
