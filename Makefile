update_dependencies:
	go mod tidy
	buf mod update

build:
	buf generate
	cd cmd/server && go build

run:
	cd cmd/server && ./server