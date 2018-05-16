package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"sync"
	"time"
)

/*
PROTOCOL the flag to switch tcp, tcp6, udp or udp6...
*/
const PROTOCOL string = "tcp"

var (
	command     string
	addr        string
	concurrency int
	count       int
	interval    int

	clientWatchGroup *sync.WaitGroup
)

func init() {
	flag.StringVar(&command, "C", "", "The command (server | client).")
	flag.StringVar(&addr, "a", ":1234", "The address.")
	flag.IntVar(&concurrency, "c", 1, "The concurrency. Client only option.")
	flag.IntVar(&count, "n", 1, "The times to be repeated. Client only option.")
	flag.IntVar(&interval, "i", 50, "The interval in milliseconds.")
	flag.Parse()
}

/*
ProcessServer the main process function for server.
*/
func ProcessServer(cc *net.TCPConn, payload []byte) {
	fmt.Println("[ProcessServer started]")

	start := time.Now().UnixNano()

	buf := make([]byte, 1024)

	count := 0

	for {
		count++

		i, err := cc.Read(buf)

		if io.EOF == err {
			break
		}

		if err != nil {
			break
		}

		// fmt.Println(string(buf[:i]))

		fmt.Printf("%d bytes received.\n", i)

		// time.Sleep(time.Duration(interval) * time.Millisecond)

		n, err := cc.Write(payload)

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("%d bytes written.\n", n)
	}
	fmt.Println("[ProcessServer stopped]", float64(time.Now().UnixNano()-start)/float64(time.Millisecond)/float64(count))
}

/*
Server The server code entrance.
*/
func Server() {
	fmt.Println("[server]")

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	laddr, err := net.ResolveTCPAddr(PROTOCOL, addr)

	if err != nil {
		panic(err)
	}

	conn, err := net.ListenTCP(PROTOCOL, laddr)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for {
		cc, err := conn.AcceptTCP()

		if err != nil {
			panic(err)
		}

		go ProcessServer(cc, bytes)
	}
}

/*
ProcessClient The main process for client.
*/
func ProcessClient(id int, raddr *net.TCPAddr, payload []byte) {
	fmt.Println("[ProcessClient started]")

	start := time.Now().UnixNano()

	conn, err := net.DialTCP(PROTOCOL, nil, raddr)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for i := 0; i < count; i++ {
		i, err := conn.Write(payload)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%d data has been written.\n", i)

		// time.Sleep(time.Duration(interval) * time.Millisecond)

		buf := make([]byte, 1024)

		n, err := conn.Read(buf)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%d data has been received.\n", n)
	}

	clientWatchGroup.Done()

	fmt.Println("[ProcessClient stopped]", float64(time.Now().UnixNano()-start)/float64(time.Millisecond)/float64(count))
}

/*
Client The client code entrance.
*/
func Client() {
	fmt.Println("[client]")

	clientWatchGroup = new(sync.WaitGroup)

	raddr, err := net.ResolveTCPAddr(PROTOCOL, addr)

	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(buf, "\x02%d\x01", len(bytes))
	buf.Write(bytes)
	buf.WriteByte('\x03')

	if concurrency > 0 {
		clientWatchGroup.Add(concurrency)

		for i := 0; i < concurrency; i++ {
			go ProcessClient(i, raddr, buf.Bytes())
		}

		clientWatchGroup.Wait()
	}
}

func main() {
	switch command {
	case "server":
		Server()
	case "client":
		Client()
	case "":
		flag.PrintDefaults()
	}
}
