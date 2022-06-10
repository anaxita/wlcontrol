-- mikrotik devices
CREATE TABLE devices (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	address TEXT NOT NULL,
	login TEXT NOT NULL,
	password TEXT NOT NULL
);

CREATE TABLE devices_chats_wl (
	chat_id INTEGER NOT NULL,
	device_id INTEGER REFERENCES devices(id) NOT NULL ,
	wl TEXT NOT NULL ,
	UNIQUE (device_id, chat_id, wl)
)