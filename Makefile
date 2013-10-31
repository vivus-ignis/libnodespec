default: deps
	go build

deps:
	go get github.com/BurntSushi/toml

install:
	go install

fmt:
	gofmt -w *.go */*.go
	colcheck *.go */*.go

tags:
	find ./ -name '*.go' -print0 | xargs -0 gotags > TAGS

push:
	git push origin master

test:
	go test -v
