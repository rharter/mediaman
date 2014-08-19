SELFPKG := github.com/rharter/mediaman
VERSION := 0.0.1
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
PKGS := \
database \
model \
api \
handler \
processor \
transcoding
PKGS := $(addprefix github.com/rharter/mediaman/pkg/,$(PKGS))
.PHONY := test $(PKGS)

all: embed build

build:
	go build -o bin/mediaman -ldflags "-X main.version $(VERSION)dev-$(SHA)" $(SELFPKG)/cmd/mediaman

build-dist: godep
	godep go build -o bin/mediaman -ldflags "-X main.version $(VERSION)-$(SHA)" $(SELFPKG)/cmd/mediaman

bump-deps: deps vendor

deps: 
	go get -u -t -v ./...

vendor: godep
	godep save ./...

# Embed static assets
embed: js rice
	cd cmd/mediaman && rice embed
	cd pkg/template && rice embed

js:
	#cd cmd/mediaman/assets && find js -name "*.js" ! -name '.*' ! -name "main.js" -exec cat {} \; > js/main.js

test:  $(PKGS)

$(PKGS):  godep
	godep go test -v $@

clean: rice
	cd cmd/mediaman && rice clean
	cd pkg/template && rice clean
	rm -rf cmd/mediaman/mediaman
	rm -rf cmd/mediaman/mediaman.sqlite
	rm -rf deb/mediaman.deb
	rm -rf bin/mediaman
	rm -rf mediaman.sqlite
	rm -rf /tmp/mediaman.sqlite
	rm -rf sample

run:
	bin/mediaman --port=":8080" --datasource="mediaman.sqlite"

godep:
	go get github.com/tools/godep

rice:
	go get github.com/GeertJohan/go.rice/rice
	go build github.com/GeertJohan/go.rice/rice

dpkg:
	mkdir -p deb/mediaman/usr/local/bin
	mkdir -p deb/mediaman/var/lib/mediaman
	mkdir -p deb/mediaman/var/cache/mediaman
	cp bin/mediaman deb/mediaman/usr/local/bin
	-dpkg-deb --build deb/mediaman

sample:
	mkdir -p "sample"
	cd sample && tar xjf ../sample-data.tar.bz2 