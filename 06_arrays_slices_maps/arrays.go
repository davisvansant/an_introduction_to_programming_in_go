package main

import "fmt"

func main()  {
  example_01()
  example_02()
}

func example_01() {
  var x [5]int
  x[4] = 100
  fmt.Println(x)
}

func example_02()  {
  x := [5]float64{
    98,
    93,
    77,
    82,
    83,
  }

  var total float64 = 0
  for _, value := range x {
    total += value
  }
  fmt.Println(total / float64(len(x)))
}
