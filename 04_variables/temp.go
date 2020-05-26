package main

import "fmt"

func main()  {
  temp := f_to_c(80)
  fmt.Println(temp)
}

func f_to_c(f float64) float64 {
  var c float64 = ((f - 32) * 5 / 9)
  return c
}
