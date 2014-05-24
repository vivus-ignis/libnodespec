#VMS := centos ubuntu arch gentoo slackware
VMS := ubuntu
GO := /usr/local/go/bin/go

default: deps
	$(GO) build

deps:
	$(GO) get github.com/BurntSushi/toml

install:
	$(GO) install

fmt:
	gofmt -w *.go */*.go
	colcheck *.go */*.go

tags:
	find ./ -name '*.go' -print0 | xargs -0 gotags > TAGS

push:
	git push origin master

start_vms:
	@for vm in $(VMS); do vagrant up $$vm; done

local_test:
	@echo "***** Local OS *****"
	@echo
	-$(GO) test -v

test: start_vms local_test
	-@for vm in $(VMS); do \
	  echo ; \
		echo "***** $$vm *****" ; \
	  vagrant provision $$vm ; \
		vagrant ssh $$vm -c 'cd /vagrant; GOPATH=/home/vagrant/go /usr/local/go/bin/go get github.com/BurntSushi/toml'; \
	  echo "Going to run unit tests..."; \
	  vagrant ssh $$vm -c 'cd /vagrant; GOPATH=/home/vagrant/go /usr/local/go/bin/go test -v' ; \
	  echo "--------------------------------------------------------------------------------" ; \
	done
