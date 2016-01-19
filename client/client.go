package client

import (
	"../protocol"
	"../util"
	"encoding/binary"
	"encoding/json"
	"log"
	"net"
)

func Run() {
	conn, err := net.Dial("tcp", "localhost:9999")
	util.CheckError(err)
	data := protocol.User{1, "u"}
	// data := protocol.Chat{"hello my friend xxxx", 1}
	sendMsg(data, conn)
}

func sendMsg(v interface{}, conn net.Conn) {
	b, err := json.Marshal(v)
	buf := make([]byte, 64)
	buf[0] = 0
	switch v.(type) {
	case protocol.User:
		buf[0] = 2
	case protocol.Chat:
		buf[0] = 1
	}
	binary.BigEndian.PutUint16(buf[1:3], uint16(len(b)))
	copy(buf[3:], b)

	_, err = conn.Write(buf)
	util.CheckError(err)
	log.Println("send msg to server >>>>>>>>")
}
