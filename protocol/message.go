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
	HeadLen uint = 1
	DataLen uint = 2
)

const (
// HeartBeat = 1
// Chat = 2
// User = 3

)

var messageTypeMap = make(map[uint8]reflect.Type)

type Chat struct {
	Message    string
	TargetUser uint
}
type User struct {
	Id   uint
	Name string
}

func init() {
	//register
	messageTypeMap[1] = reflect.TypeOf(Chat{})
	messageTypeMap[2] = reflect.TypeOf(User{})
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
	headByte := make([]byte, HeadLen)
	_, err := reader.Read(headByte)
	util.CheckError(err)
	messageType := uint8(headByte[0])
	for key, value := range messageTypeMap {
		if messageType == key {
			return reflect.New(value).Interface()
		}
	}
	panic("undefined msg type")
}

func parseMessageData(reader io.Reader) []byte {
	lenByte := make([]byte, DataLen)
	_, err := reader.Read(lenByte)
	util.CheckError(err)
	messageLen := binary.BigEndian.Uint16(lenByte)
	messageByte := make([]byte, messageLen)
	_, err = reader.Read(messageByte)
	util.CheckError(err)
	return messageByte
}
