build-prod:
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui" -o kmsbot.exe .

migrate-install:
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate-new:
	migrate create -ext sql -dir migrations "$(name)"

migrate-up:
	migrate -database "sqlite3://bot.db" -path migrations up

migrate-down:
	migrate -database "sqlite3://bot.db?" -path migrations down 1