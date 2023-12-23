build:
	go build -o ./bin/gpsql ./cmd/main.go
run:
	go run ./cmd/main.go
down:
	go run ./cmd/main.go down
up:
	go run ./cmd/main.go up
