.DEFAULT_GOAL := run

include: .env

run:
	go run main.go

first_run:
	cp .env.example .env && go run main.go