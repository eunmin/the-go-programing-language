package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		go showClock("ok", arg)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func showClock(tz string, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
