package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

const PROTOCOL string = "tcp"

var (
	command string
	addr    string
)

func Init() {
	flag.StringVar(&command, "c", "server", "The command (server | client).")
	flag.StringVar(&addr, "a", ":1234", "The address.")
	flag.Parse()

	fmt.Printf("command : %s, addr : %s \n", command, addr)
}

func Process(cc *net.TCPConn) {
	buf := make([]byte, 1024)

	for {
		i, err := cc.Read(buf)

		if io.EOF == err {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(string(buf[:i]))
	}
}

func Server() {
	fmt.Println("[server]")
	laddr, err := net.ResolveTCPAddr(PROTOCOL, addr)

	if err != nil {
		panic(err)
	}

	conn, err := net.ListenTCP(PROTOCOL, laddr)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	cc, err := conn.AcceptTCP()

	if err != nil {
		panic(err)
	}

	Process(cc)
}

func Client() {
	fmt.Println("[client]")

	raddr, err := net.ResolveTCPAddr(PROTOCOL, addr)

	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP(PROTOCOL, nil, raddr)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	str := string(bytes)

	o := fmt.Sprintf("\x02%d\x01%s\x03", len(bytes), str)

	i, err := conn.Write([]byte(o))

	if err != nil {
		panic(err)
	}

	fmt.Println("i : ", i)

	x, err := io.Copy(os.Stdout, conn)

	fmt.Println("x : ", x)
}

func main() {
	Init()

	switch command {
	case "server":
		Server()
	case "client":
		Client()
	}
}
