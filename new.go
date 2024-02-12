package main
import "fmt"
func main() {
    a, b := 2, 3
    c := float64(a + b)
    fmt.Printf("%v + %v = %f = %v, stored as %T", a, b, c, c, c)
}