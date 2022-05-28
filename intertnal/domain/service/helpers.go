package delivery

import "log"

func catchPanic() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("[PANIC] ", err)
		}
	}()
}
