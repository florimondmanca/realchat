INPUT=cmd/server/main.go
OUTPUT=cmd/server/main.out

all: build test

install:
	go install

build:
	go build -o ${OUTPUT} ${INPUT}

test:
	go test -v ${INPUT}

run:
	make build
	./${OUTPUT}

check:
	go vet

fmt:
	gofmt -w ./$*

clean:
	go clean
	rm ${OUTPUT}
