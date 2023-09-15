package main
 
import "fmt"
 
func main() {
  var n int
  fmt.Scan(&n)
  if n == 15 || n == 20 || n == 21 {
    fmt.Println("NO")
  } else {
    fmt.Println("YES")
  }
}
