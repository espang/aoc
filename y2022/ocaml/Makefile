.PHONY: clean

day?=day1

clean:
	rm -rf _build

build:
	dune build main.exe

run: build
	./_build/default/main.exe ${day}
