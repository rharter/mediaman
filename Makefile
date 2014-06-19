SELFPKG := github.com/rharter/mediaman
VERSION := 0.0.1
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
PKGS := \
database \
model
PKGS := $(addprefix github.com/drone/drone/pkg/,$(PKGS))
.PHONY := test $(PKGS)

all: embed build

build:
	go build -o bin/mediaman -ldflags "-X main.version $(VERSION)dev-$(SHA)" $(SELFPKG)/cmd/mediaman

build-dist:
	godep go build -o bin/mediaman -ldflags "-X main.version $(VERSION)-$(SHA)" $(SELFPKG)/vmd/mediaman

bump-deps: deps

deps:
	go get -u -t -v ./...

# Embed static assets
embed: js rice
	cd cmd/mediaman && rice embed
	cd pkg/template && rice embed

js:
	cd cmd/mediaman/assets && find js -name "*.js" ! -name '.*' ! -name "main.js" -exec cat {} \; > js/main.js

test: $(PKGS)

$(PKGS): godep
	godep go text -v $@

clean: rice
	cd cmd/mediaman && rice clean
	cd pkg/template && rice clean
	rm -rf cmd/mediaman/mediaman
	rm -rf cmd/mediaman/mediaman.sqlite
	rm -rf bin/mediaman
	rm -rf mediaman.sqlite
	rm -rf /tmp/mediaman.sqlite

run:
	bin/mediaman --port=":8080" --datasource="mediaman.sqlite"

godep:
	go get github.com/tools/godep

rice:
	go install github.com/GeertJohan/go.rice/rice