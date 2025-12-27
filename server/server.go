package server

import (
	"bufio"
	"net"
	"strings"

	"Minireddis/command"
	"Minireddis/storage"
)

func StartTCP(addr string, store *storage.Store) {
	ln, _ := net.Listen("tcp", addr)
	println("🚀 TCP server on", addr)

	for {
		conn, _ := ln.Accept()
		go handle(conn, store)
	}
}

func handle(conn net.Conn, store *storage.Store) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		resp := command.HandleCommand(line, store)
		if resp != "" {
			conn.Write([]byte(resp + "\n"))
		}
	}
}
