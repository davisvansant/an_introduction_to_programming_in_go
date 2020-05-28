package main

import "fmt"

func main() {
  nextOdd := makeOddGenerator()

  for number := 0; number <= 10; number++ {
    fmt.Println(nextOdd())
  }
}

func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
