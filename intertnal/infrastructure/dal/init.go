package dal

import (
	"fmt"
	"github.com/gocraft/dbr"
	"time"
	"wlcontrol/intertnal/domain"
	"wlcontrol/intertnal/entity"
	"wlcontrol/intertnal/infrastructure/dal/cache"
)

type Repository struct {
	db    *dbr.Connection
	cache cache.Cache
}

func NewRepository(dbName string) (*Repository, error) {
	conn, err := dbr.Open("sqlite3",
		fmt.Sprintf("file:%s.db", dbName), nil)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	return &Repository{
		db:    conn,
		cache: cache.New(),
	}, nil
}

func (r *Repository) ChatUserState(chatID, userID int64) (entity.User, error) {
	return r.cache.User(chatID, userID)
}

func (r *Repository) AddChatUserState(chatID int64, u entity.User) {
	r.cache.AddUser(chatID, u)
}

func (r *Repository) DeleteChatUserState(chatID, userID int64) {
	r.cache.DeleteUser(chatID, userID)
}

func (r *Repository) AddRouter(router entity.MikrotikCreate) error {
	_, err := r.db.NewSession(nil).
		InsertInto("devices").
		Columns("name", "address", "login", "password").
		Record(router).
		Exec()

	return err
}

func (r *Repository) ChatByID(id int64) (entity.Chat, error) {
	var devices []entity.Mikrotik

	_, err := r.db.NewSession(nil).
		Select("id, name, address, login, password, chat_id, wl").
		From("devices_chats_wl").
		Join("devices", "devices.id = devices_chats_wl.device_id").
		Where("chat_id = ?", id).
		GroupBy("id").
		Load(&devices)
	if err != nil {
		return entity.Chat{}, err
	}

	if len(devices) == 0 {
		return entity.Chat{}, domain.ErrNotFound
	}

	return entity.Chat{ID: devices[0].ChatID, Devices: devices}, nil
}

func (r *Repository) AddDevicesToChat(devices ...entity.Mikrotik) error {
	stmt, err := r.db.NewSession(nil).Prepare("INSERT INTO devices_chats_wl (chat_id, device_id, wl) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range devices {
		_, err = stmt.Exec(v.ChatID, v.ID, v.WL)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) RemoveDeviceFromChat(devices ...entity.Mikrotik) error {
	stmt, err := r.db.NewSession(nil).Prepare("DELETE FROM devices_chats_wl WHERE chat_id = ? AND device_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range devices {
		_, err = stmt.Exec(v.ChatID, v.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) DeviceByID(id int64) (device entity.Mikrotik, err error) {
	err = r.db.NewSession(nil).
		Select("id, name, address, login, password").
		From("devices").
		Where("id = ?", id).
		LoadOne(&device)

	return
}

func (r *Repository) Devices() (devices []entity.Mikrotik, err error) {
	_, err = r.db.NewSession(nil).
		Select("id, name, address, login, password").
		From("devices").
		Load(&devices)

	return
}
