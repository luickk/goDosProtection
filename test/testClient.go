package main

import (
  "fmt"
  "net"
)

func main() {
  conn, err := net.Dial("tcp", "localhost:8081")
  if err != nil {
    fmt.Println(err)
    return
  }

  conn.Close()
}
