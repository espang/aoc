-include .env

export AOC_ROOT_FOLDER = $(shell pwd)

.PHONY: y2016 y2022

y2016:
	go run y2016/main.go ${AOC_DAY}

y2022:
	go run y2022/main.go ${AOC_DAY}

test:
	go test ./...