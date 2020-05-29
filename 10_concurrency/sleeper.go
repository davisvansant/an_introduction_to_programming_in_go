package main

import (
  "fmt"
  "time"
)

func sleep(x int) {
  for i := 0; i < 5; i++ {
    select {
    case <-time.After(time.Second * time.Duration(x)):
      fmt.Println("I'm sleeping for", x, "seconds")
    }
  }
}

func main()  {
  go sleep(3)
  var input string
  fmt.Scanln(&input)
}
