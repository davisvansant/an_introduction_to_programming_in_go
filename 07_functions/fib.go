package main

import "fmt"

func main()  {
  for number := uint(0); number <= 40; number++ {
    fmt.Println(fib(number))
  }
}

func fib(n uint) uint {
  if n <= 1 {
    return n
  }
  return fib(n - 1) + fib(n - 2)
}
