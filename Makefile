# variable
# SRCS=$(shell find . -type f -name "*.go")
# environment
.DEFAULT_GOAL := all

all:
	make build/api
	make build/web
	make run

build/api:
	cd api && go build main.go

build/web:
	echo hoge

run:
	cd api && ./main