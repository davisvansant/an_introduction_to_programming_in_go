package main

import (
  "fmt"
  "crypto/sha1"
)

// sha1 is broken and shouldnt be used

func main()  {
  h := sha1.New()
  h.Write([]byte("tester"))
  bs := h.Sum([]byte{})
  fmt.Println(bs)
}
