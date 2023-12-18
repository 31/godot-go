.DEFAULT_GOAL := build

GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
GODOT?=$(shell which godot)
CWD=$(shell pwd)

OUTPUT_PATH=test/demo/lib
CGO_ENABLED=1

.PHONY: goenv installdeps generate update_godot_headers_from_binary build clean_src clean remote_debug_test test interactive_test open_demo_in_editor

goenv:
	go env

installdeps:
	go install golang.org/x/tools/cmd/goimports@latest

generate: installdeps clean
	go generate
	go run ./internal/cmd/generateutil

update_godot_headers_from_binary: ## update godot_headers from the godot binary
	go run ./internal/cmd/updategodotheadersfrombinary

build: goenv
	go env
	go run ./internal/cmd/build

clean_src:
	go run ./internal/cmd/cleansrc

clean: clean_src
	go run ./internal/cmd/clean

remote_debug_test:
	CI=1 \
	LOG_LEVEL=info \
	GOTRACEBACK=crash \
	GODEBUG=sbrk=1,asyncpreemptoff=1,cgocheck=0,invalidptr=1,clobberfree=1,tracebackancestors=5 \
	gdbserver --once :55555 $(GODOT) --headless --verbose --debug --path test/demo/

ci_gen_test_project_files:
	go run ./internal/cmd/cigentestprojectfiles

test:
	go run ./internal/cmd/test

interactive_test:
	go run ./internal/cmd/interactivetest

open_demo_in_editor:
	go run ./internal/cmd/opendemoineditor
