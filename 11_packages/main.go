package main

import "fmt"
import "./math"

func main() {
  xs := []float64{1, 2, 3, 4}
  avg := math.Average(xs)
  fmt.Println(avg)

  nums := []float64{10, 20, 12, 92340, 99, 83423}
  fmt.Println("The largest is", math.Max(nums))
  fmt.Println("The smallest is", math.Min(nums))

  nothing := []float64{}
  fmt.Println(math.Average(nothing))
}
