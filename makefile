.PHONY: all fmt tags doc

all:
	go install -v ./...

rall:
	go build -a ./...

fmt:
	gofmt -s -w -l .

tags:
	gotags `find . -name "*.go"` > tags

test:
	go test ./...

testv:
	go test -v ./...

lc:
	wc -l `find . -name "*.go"` 

doc:
	godoc -http=:8000

lint:
	golint `find . -name "*.go"`
