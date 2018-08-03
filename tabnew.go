package main

import (
  "os"
  "fmt"
  "net"
  "path/filepath"
  "github.com/msgpack-rpc/msgpack-rpc-go/rpc"
)

func main() {
  conn, err := net.Dial("unix", os.Getenv("NVIM_LISTEN_ADDRESS"))
  if err != nil {
    fmt.Println("fail to connect to server1")
    return
  }
  client := rpc.NewSession(conn, true)
  if len(os.Args) <= 1 {
    client.Send("nvim_command", "tabnew")
  }else{
    path, _ := filepath.Abs(os.Args[1])
    client.Send("nvim_command", "tabnew " + path)
  }
  pwd, _ := filepath.Abs(".")
  client.Send("nvim_command", "lcd " + pwd)
}
