package main

import "fmt"

func main()  {
  fmt.Println(factorial(1))
  fmt.Println(factorial(10))
  fmt.Println(factorial(20))
}

func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  return x * factorial(x - 1)
}
