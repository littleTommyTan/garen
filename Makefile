# Go Makefile

all: run

test:
	go test -v ./

run:
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -o dist/garen

deps:
	go get -u github.com/kardianos/govendor
	govendor sync
	git clone https://github.com/golang/crypto.git vendor/golang.org/x/crypto/
	git clone https://github.com/golang/net.git vendor/golang.org/x/net/
	git clone https://github.com/golang/sys.git vendor/golang.org/x/sys/
	git clone https://github.com/golang/time.git vendor/golang.org/x/time/
	git clone https://github.com/golang/text.git vendor/golang.org/x/text/
	git clone https://github.com/grpc/grpc-go.git vendor/google.golang.org/grpc
	git clone https://github.com/google/go-genproto.git vendor/google.golang.org/genproto
	git clone https://github.com/google/protobuf.git vendor/google.golang.org/protobuf
