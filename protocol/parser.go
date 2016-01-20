package protocol

import (
	"../util"
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"reflect"
)

// ======================
// | HEAD | LENGTH | DATA
// ======================
const (
	headLen uint = 1
	dataLen uint = 2
)

var messageTypeMap = make(map[uint8]reflect.Type)

func registerMessage(msgType uint8, msg interface{}) {
	messageTypeMap[msgType] = reflect.TypeOf(msg)
}

func ParseMessage(reader io.Reader) interface{} {
	v := parseMessageType(reader)
	data := parseMessageData(reader)
	err := json.Unmarshal(data, &v)
	util.CheckError(err)
	log.Println(v)
	return v
}

func parseMessageType(reader io.Reader) interface{} {
	headByte := make([]byte, headLen)
	_, err := reader.Read(headByte)
	util.CheckError(err)
	messageType := uint8(headByte[0])
	log.Println("messageType:", messageType)
	for key, value := range messageTypeMap {
		if messageType == key {
			return reflect.New(value).Interface()
		}
	}
	panic("undefined msg type")
}

func parseMessageData(reader io.Reader) []byte {
	lenByte := make([]byte, dataLen)
	_, err := reader.Read(lenByte)
	util.CheckError(err)
	messageLen := binary.LittleEndian.Uint16(lenByte)
	log.Println("messageLen:", messageLen)
	messageByte := make([]byte, messageLen)
	_, err = reader.Read(messageByte)
	util.CheckError(err)
	return messageByte
}
