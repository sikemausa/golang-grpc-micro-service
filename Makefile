build:
	buf generate
	cd cmd/server && go build

run:
	cd cmd/server && ./server