package entity

type User struct {
	ID        int64
	MessageID int
	State     UserState
}

type UserState uint8

const (
	UserStateDefault = iota
	UserStateAddRouter
)
