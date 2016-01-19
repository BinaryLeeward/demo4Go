package server

import (
	"../protocol"
	"../util"
	"log"
	"net"
	"reflect"
)

func Run() {
	l, err := net.Listen("tcp", ":9999")
	util.CheckError(err)
	defer l.Close()
	for {
		conn, err := l.Accept()
		util.CheckError(err)
		log.Println("receive msg from client <<<<<<<<<<<<<<<")
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	v := protocol.ParseMessage(conn)
	log.Println(reflect.TypeOf(v))
	switch v.(type) {
	case protocol.User:
		x := v.(protocol.User)
		log.Println(x.Name)
	}
}
