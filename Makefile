-include .env

VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOPATH)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies.
install: go-get

go-get: 
	go get -u ./...

## start: start dota item/piece api/server and run discord bot
start:
	@echo " > running bot & server"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run main.go -port "8080" & 
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run discord_bot.go -t $(shell echo; read -p "Token: " tkn; echo $$tkn)


.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
