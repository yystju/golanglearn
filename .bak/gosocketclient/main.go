package main

import (
	"io"
	//	"log"
	"net"
	"os"
)

func process(a *TCPAddr, data []byte) {
	c, err := net.DialTCP("tcp", nil, addr)

	if err != nil {
		panic(err)
	}

	go io.Copy(os.Stdout, c)
	io.Copy(c, os.Stdin)
}

func main() {
	server := os.Args[1]

	addr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		panic(err)
	}
}
