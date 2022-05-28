package dal

import (
	"wlcontrol/intertnal/domain/entity"
	"wlcontrol/intertnal/infrastructure/dal/cache"
)

type Repository struct {
	cache cache.Cache
}

func NewRepository() Repository {
	return Repository{
		cache: cache.New(),
	}
}

func (r *Repository) ChatUser(chatID, userID int64) (entity.User, error) {
	return r.cache.ChatUser(chatID, userID)
}

func (r *Repository) AddChatUser(chatID int64, u entity.User) {
	r.cache.AddChatUser(chatID, u)
}

func (r *Repository) DeleteChatUser(chatID, userID int64) {
	r.cache.DeleteChatUser(chatID, userID)
}
