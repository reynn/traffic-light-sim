APP_NAME := traffic-lights

.PHONY: build
build:
	go build -race -o $(APP_NAME) -ldflags "-w -s" -trimpath cmd/traffic/main.go

.PHONY: test
test:
	go test -race ./...

.PHONY: run
run: build
	./$(APP_NAME)

.PHONY: clean
clean:
	rm -f ./$(APP_NAME)
