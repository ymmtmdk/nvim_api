package main

import (
	"fmt"
	"github.com/msgpack-rpc/msgpack-rpc-go/rpc"
	"net"
	"os"
	// "path/filepath"
	"strings"
)

func main() {
	conn, err := net.Dial("unix", os.Getenv("NVIM_LISTEN_ADDRESS"))
	if err != nil {
		fmt.Println("fail to connect to server")
		return
	}
	// fmt.Println(filepath.Abs(os.Args[2]))
	msg := strings.Join(os.Args[2:], " ")
	client := rpc.NewSession(conn, true)
	fmt.Println(msg)

	retval, xerr := client.Send(os.Args[1])
	// retval, xerr := client.Send("nvim_command", "echo 1")

	if xerr != nil {
		fmt.Println(xerr)
		return
	}
	fmt.Println(retval)
}
