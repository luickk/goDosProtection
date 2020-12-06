package main

import (
  "goDosProtection/pkg/goDosProtection"
)

func main() {
  dp := goDosProtection.New(5)

  ln, err := net.Listen("tcp", "127.0.0.1:8081")
   if err != nil {
       log.Fatal(err)
   }

  for {
     conn, err := ln.Accept()
     if err != nil {
         log.Fatal(err)
     }
     // not banned
     if !dp.Client(conn) {
       handleConn(conn)
    } else {

    }
  }
}

func handleConn(conn net.Conn) {

}
