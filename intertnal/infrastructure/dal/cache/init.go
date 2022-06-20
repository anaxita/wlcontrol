package cache

import (
	"sync"
	"wlcontrol/intertnal/domain"
	"wlcontrol/intertnal/entity"
)

type Cache struct {
	m sync.RWMutex

	// users // UserID -> ChatID -> User
	users map[int64]map[int64]entity.User
}

func New() Cache {
	return Cache{
		users: make(map[int64]map[int64]entity.User),
	}
}

func (c *Cache) User(chatID, userID int64) (entity.User, error) {
	c.m.RLock()
	defer c.m.RUnlock()

	chats, ok := c.users[chatID]
	if !ok {
		return entity.User{}, domain.ErrNotFound
	}

	u, ok := chats[userID]
	if !ok {
		return u, domain.ErrNotFound
	}

	return u, nil
}

func (c *Cache) AddUser(chatID int64, u entity.User) {
	c.m.Lock()
	defer c.m.Unlock()

	_, ok := c.users[chatID]
	if !ok {
		c.users[chatID] = make(map[int64]entity.User)
	}

	c.users[chatID][u.ID] = u
}

func (c *Cache) DeleteUser(chatID, userID int64) {
	c.m.Lock()
	defer c.m.Unlock()

	chats, ok := c.users[chatID]
	if ok {
		delete(chats, userID)
	}
}
