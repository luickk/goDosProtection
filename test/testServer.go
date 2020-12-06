package main

import (
  "goDosProtection"
  "net"
  "strings"
  "fmt"
)

func main() {
  // initiating DOS protection with 10 second reconnection delay
  dp := goDosProtection.New(10)

  ln, err := net.Listen("tcp", "127.0.0.1:8081")
   if err != nil {
       fmt.Println(err)
       return
   }

  for {
     conn, err := ln.Accept()
     if err != nil {
         fmt.Println(err)
         return
     }
     // not banned
     if !dp.Client(strings.Split(conn.RemoteAddr().String(), ":")[0]) {
       fmt.Println("accepted client conn")
       handleConn(conn)
    } else {
      fmt.Println("refused connection")
    }
  }
}

func handleConn(conn net.Conn) {

}
