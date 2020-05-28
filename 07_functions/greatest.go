package main

import "fmt"

func main()  {
  numbers := []int{ 100, 23, 5082, 2348, 96, 8532, 320032942, 12093123085308, 20390092 }
  fmt.Println(greatest(numbers...))
}

func greatest(nums ...int) int {
  greatest := 0
  for _, n := range nums {
    if n > greatest {
      greatest = n
    }
  }
  return greatest
}
