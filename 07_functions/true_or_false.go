package main

import "fmt"

func true_or_false(integer int) bool {
  // half := integer / 2
  if integer % 2 == 0 {
    return true
  } else {
    return false
  }
}

func main()  {
  fmt.Println(true_or_false(1))
  fmt.Println(true_or_false(2))
  fmt.Println(true_or_false(37))
  fmt.Println(true_or_false(8080))
}
