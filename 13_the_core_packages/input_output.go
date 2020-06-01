package main

import (
  "bytes"
  "fmt"
  "io"
  "log"
  "strings"
)

func main()  {
  var buf bytes.Buffer
  fmt.Println(buf.Write([]byte("test")))
  fmt.Println(buf.Bytes())
  if _, err := io.Copy(&buf, strings.NewReader("cool")); err != nil {
    log.Fatal(err)
  }
  fmt.Println(buf.String())
}
