package remoteserver

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

type StringServer string

func (s StringServer) ServerHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		
	}
}

const addr = "localhost:7070"

func TestExec(t *testing.T) {
	s := createServer(addr)
	go s.ListenAndServe()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHost:localhost:7070\r\n\r\n")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(conn)
	conn.SetReadDeadline(time.Now().Add(time.Second))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	ctx, tt := context.WithTimeout(context.Background(), 5*time.Second)
	if tt != nil {
		ctx.Deadline()
	}
	s.Shutdown(ctx)
}
