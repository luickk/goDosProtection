# Go DOS Protection

Simple and very low level ip-address(string) based server side denial of service attack protection by limiting connection frequency.

Server Example:
``` go
// initiating DOS protection with 10 second reconnection delay
dp := goDosProtection.New(10)

conn, err := ln.Accept()
if err != nil {
    fmt.Println(err)
    return
}
// client is not banned
if !dp.Client(strings.Split(conn.RemoteAddr().String(), ":")[0]) {
  fmt.Println("Accepted client connection")
  handleConn(conn)
// client is banned
} else {
 fmt.Println("Refused client connection")
}
```
