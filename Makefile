ARCH=$(shell arch)
UNAME=$(shell uname)
SHORTHASH=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)-$(SHORTHASH)
FILEVERSION=$(VERSION)-$(UNAME)-$(ARCH)

all: gomrefdash

.PHONY: version
version:
	echo $(VERSION)

gomrefdash:
	GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -o gomrefdash -ldflags "-X main.dver=$(VERSION)"

.PHONY: package
package: gomrefdash gomrefdash-$(FILEVERSION).tar.gz

gomrefdash-$(FILEVERSION).tar.gz: gomrefdash
	tar cfz gomrefdash-$(FILEVERSION).tar.gz gomrefdash templates static .env.sample

.PHONY: docker
docker:
	docker build -t gomrefdash .

.PHONY: docker-clean
docker-clean:
	-docker image rm gomrefdash

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose up -d
	docker-compose logs -f

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose down

.PHONY: docker-compose-logs
	docker-compose logs -f

.PHONY: run
run:
	go run -ldflags "-X main.dver=$(VERSION)" *.go

.PHONY: clean
clean:
	-rm gomrefdash
	-rm -f gomrefdash-$(FILEVERSION).tar.gz