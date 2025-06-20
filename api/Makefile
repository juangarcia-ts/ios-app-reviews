#!make
include .env
export

install:
	go mod tidy

run:
	go run cmd/api/main.go

migration/new:
	sql-migrate new -config=sqlmigrate.yml -env="default" $(name)

migration/up:
	sql-migrate up -config=sqlmigrate.yml -env="default"

migration/down:
	sql-migrate down -config=sqlmigrate.yml -env="default"