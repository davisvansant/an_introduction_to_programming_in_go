package main

import "fmt"

func sum(nums ...int) int {
  total := 0
  for _, n := range nums {
    total += n
  }
  return total
}

func main()  {
  numbers := []int{10, 20, 30, 40, 50}

  fmt.Println(sum(numbers...))
}
