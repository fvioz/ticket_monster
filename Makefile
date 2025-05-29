deps:
	brew bundle

build:
	cd api && go get -u github.com/swaggo/swag/cmd/swag && go build -o bin/api main.go && $(HOME)/go/bin/swag init
	cd fetcher && go build -o bin/api main.go
	cd processors && go build -o bin/api main.go

run:
	podman-compose -f docker-compose.dev.yml up

.PHONY: clean
clean:
	podman-compose -f docker-compose.dev.yml down
