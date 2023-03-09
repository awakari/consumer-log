.PHONY: test clean
default: build

BINARY_FILE_NAME=consumer-log

proto:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	PATH=${PATH}:~/go/bin protoc \
		--go_out=plugins=grpc:. \
		--go_opt=paths=source_relative \
		api/grpc/*.proto \
		api/grpc/queue/*.proto

vet: proto
	go vet

test: vet
	go test -race ./...

build: proto
	CGO_ENABLED=0 GOOS=linux GOARCH= GOARM= go build -o ${BINARY_FILE_NAME} main.go
	chmod ugo+x ${BINARY_FILE_NAME}

docker:
	docker build -t awakari/consumer-log .

run: docker
	docker run \
		-d \
		--name awakari-consumer-log \
		-p 8080:8080 \
		--expose 8080 \
		awakari/consumer-log

staging: docker
	./scripts/staging.sh

release: docker
	./scripts/release.sh

clean:
	go clean
	rm -f ${BINARY_FILE_NAME}
