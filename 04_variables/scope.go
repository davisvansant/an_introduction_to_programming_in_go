package main

import "fmt"

var x string = "Hello World, from scope"

func main() {
  fmt.Println(x)
  f()
}

func f() {
  fmt.Println(x)
}
