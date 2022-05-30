package entity

type User struct {
	ID        int64
	ChatID    int64
	MessageID int

	MikrotikID int64
	State      UserState
}

type UserState uint8

const (
	UserStateAddRouter UserState = iota + 1
	UserStateEnterChatID
	UserStateEditChatSettings
	UserStateSetDeviceToChat
)
