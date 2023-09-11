-include .env
export AOC_ROOT_FOLDER = $(shell pwd)

.PHONY: y2016 y2017 y2022 help test
.DEFAULT_GOAL := help

## runs the year 2016 go solution. Select the day by setting the `AOC_DAY` env var, e.g. `AOC_DAY=day1 make y2016`
y2016:
	go run y2016/main.go ${AOC_DAY}

## runs the year 2016 go solution. Select the day by setting the `AOC_DAY` env var, e.g. `AOC_DAY=day1 make y2017`
y2017:
	go run y2017/main.go ${AOC_DAY}


## runs the year 2022 go solution. Select the day by setting the `AOC_DAY` env var.
y2022:
	go run y2022/main.go ${AOC_DAY}

## runs the go tests
test:
	go test -count=1 -v ./...

## print help
help:
	@printf "\nusage : make <commands> \n\nthe following commands are available : \n\n"
	@cat Makefile | awk '1;/help:/{exit}' | awk '/##/ { print; getline; print; }' | awk '{ getline x; print x; }1' | awk '{ key=$$0; getline; print key "\t\t\t\t " $$0;}'
	@printf "\n"