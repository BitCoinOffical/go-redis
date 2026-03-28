#имя/цель: зависимость(может быть пустой)
# 	набор команд

APP = .bin/app.exe 
CLI = .bin/cli.exe

GOOS ?= windows


.PHONY: build

command:
	cmd /c echo $(origin GOOS)

$(APP): app/main.go
	go build -o $@ $^

$(CLI): cli/main.go
	go build -o $@ $^

clear:
ifeq ($(GOOS), windows)
	cmd /c rmdir /S /Q .bin
else
	rm -rf .bin
endif

run-cli:
	go run cli/main.go

run-app:
	go run app/main.go

all: $(APP) $(CLI)