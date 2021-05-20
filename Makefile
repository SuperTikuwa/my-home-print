# variable
api:= ./api/main.go

# environment
.DEFAULT_GOAL := all

all:
	make build/api
	make build/web

build/api:
	go build $(api)

build/web:
	echo hoge