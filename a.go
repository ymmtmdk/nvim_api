package main

import (
	"fmt"
	"github.com/neovim/go-client/nvim"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("unix", os.Getenv("NVIM_LISTEN_ADDRESS"))
	if err != nil {
		fmt.Println("fail to connect to server")
		return
	}
	nv, err := nvim.New(conn, conn, conn, log.Printf)
	if err != nil {
		fmt.Println("fail to connect to server")
		return
	}
	fmt.Println(nv)
	// var result string
	// nv.Command("echo", "1")
	// var result string
	b := nv.NewBatch()
	b.Command("echo 2")
	b.Execute()
	fmt.Println("ok")
}
