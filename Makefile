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

materials/rubygem-rake_10.3.2_all.deb:
	vagrant ssh ubuntu -c 'cd /vagrant/materials; fpm -s gem -t deb -v 10.3.2 rake'

build_test_packages: materials/rubygem-rake_10.3.2_all.deb

local_test:
	@echo "***** Local OS *****"
	@echo
	-$(GO) test -v

#test: build_test_packages start_vms local_test
test: build_test_packages start_vms
	-@for vm in $(VMS); do \
	  echo ; \
		echo "***** $$vm *****" ; \
	  vagrant provision $$vm ; \
		vagrant ssh $$vm -c 'cd /vagrant; GOPATH=/home/vagrant/go /usr/local/go/bin/go get github.com/BurntSushi/toml'; \
	  echo "Going to run unit tests..."; \
	  vagrant ssh $$vm -c 'cd /vagrant; GOPATH=/home/vagrant/go sudo -E /usr/local/go/bin/go test -v' ; \
	  echo "--------------------------------------------------------------------------------" ; \
	done
