package main

import "fmt"

func swap(x *int, y *int) {
  old_x := *x
  old_y := *y
  *x = old_y
  *y = old_x
}

func main()  {
  x := 1
  y := 2
  fmt.Println("x was", x)
  fmt.Println("y was", y)
  swap(&x, &y)
  fmt.Println("x is now", x)
  fmt.Println("y is now", y)
}
