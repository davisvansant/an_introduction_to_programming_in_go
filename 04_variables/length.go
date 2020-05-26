package main

import "fmt"

func main()  {
  length := f_to_m(300)
  fmt.Println(length)
}

func f_to_m(f float64) float64 {
  const foot float64 = 0.3048
  var m float64 = foot * f
  return m
}
