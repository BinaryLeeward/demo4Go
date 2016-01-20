package protocol

//msg type
const (
	chat = 1
	user = 2
)

type Chat struct {
	Content    string
	TargetUser uint
}
type User struct {
	Id   uint
	Name string
}

func init() {
	registerMessage(chat, Chat{})
	registerMessage(user, User{})
}
