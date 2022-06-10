package entity

type User struct {
	ID           int64     // user id
	ChatID       int64     // message from id
	MessageID    int       // message id
	EditedChatID int64     // chat to change
	MikrotikID   int64     // device id to change
	State        UserState // user start
}

type UserState uint8

const (
	UserStateAddRouter UserState = iota + 1
	UserStateEnterChatID
	UserStateEditChatSettings
	UserStateSetDeviceToChat
)
