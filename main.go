package main

import "fmt"
import "os"

import "strings"

import "net"
import "bufio"

import "wrapper"

func dumpEnv() {
	fmt.Println(os.Args)

	x, y, xy := tripple(3, 5)

	fmt.Println(x, y, xy)

	fmt.Println(os.Getenv("KEY1"))

	for i, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(i, pair[0], pair[1])
	}
}

func tripple(x int, y int) (int, int, int) {
	return x, y, x + y
}

func doServer() {
	addr, _ := net.LookupTXT("www.baidu.com")

	fmt.Println(addr)

	ln, err := net.Listen("tcp", ":1234")

	if err != nil {
		fmt.Println(err)
	}

	conn, _ := ln.Accept()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}

func main() {
	n, str := wrapper.LoadFromStdIn()
	fmt.Printf("\x02%d\x01%s\x03\n", n, str)
}
