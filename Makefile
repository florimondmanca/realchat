.PHONY: client

INPUT=cmd/server/main.go
TEST_INPUT=./pkg/app
OUTPUT=cmd/server/main.out
YARN=yarn --cwd client

all: build test

install:
	go install ${INPUT}
	${YARN} install

build-server:
	go build -o ${OUTPUT} ${INPUT}

build-client:
	${YARN} build

build:
	make build-server
	make build-client

test:
	go test -v ${TEST_INPUT}

serve:
	make build-server
	./${OUTPUT}

client:
	${YARN} start

client-dev:
	${YARN} dev

check:
	go vet ${INPUT}

fmt:
	gofmt -w ./$*

clean:
	go clean
	rm ${OUTPUT}
