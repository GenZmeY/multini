NAME     = multini
VERSION := $(shell git describe)
GOCMD    = go
LDFLAGS := "$(LDFLAGS) -s -w -X 'main.Version=$(VERSION)'"
GOBUILD  = $(GOCMD) build -ldflags=$(LDFLAGS)
SRCMAIN  = ./cmd/$(NAME)
SRCDOC   = ./doc
SRCTEST  = ./test
BINDIR   = bin
BIN      = $(BINDIR)/$(NAME)
README   = $(SRCDOC)/README
LICENSE  = LICENSE
TEST     = $(SRCTEST)/run_test.sh
PREFIX   = /usr

.PHONY: all prep doc build check-build freebsd-386 darwin-386 linux-386 windows-386 freebsd-amd64 darwin-amd64 linux-amd64 windows-amd64 compile install check-install uninstall clean test

all: build

prep: clean
	go mod init; go mod tidy
	mkdir $(BINDIR)
	
doc: check-build
	test -d $(SRCDOC) || mkdir $(SRCDOC)
	$(BIN) --help > $(README)
	
build: prep
	$(GOBUILD) -o $(BIN) $(SRCMAIN)

check-build:
	test -e $(BIN)

freebsd-386: prep
	GOOS=freebsd GOARCH=386 $(GOBUILD) -o $(BIN)-freebsd-386 $(SRCMAIN)

darwin-386: prep
	GOOS=darwin GOARCH=386 $(GOBUILD) -o $(BIN)-darwin-386 $(SRCMAIN)

linux-386: prep
	GOOS=linux GOARCH=386 $(GOBUILD) -o $(BIN)-linux-386 $(SRCMAIN)

windows-386: prep
	GOOS=windows GOARCH=386 $(GOBUILD) -o $(BIN)-windows-386.exe $(SRCMAIN)
	
freebsd-amd64: prep
	GOOS=freebsd GOARCH=amd64 $(GOBUILD) -o $(BIN)-freebsd-amd64 $(SRCMAIN)

darwin-amd64: prep
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BIN)-darwin-amd64 $(SRCMAIN)

linux-amd64: prep
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN)-linux-amd64 $(SRCMAIN)

windows-amd64: prep
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BIN)-windows-amd64.exe $(SRCMAIN)

compile: freebsd-386 darwin-386 linux-386 windows-386 freebsd-amd64 darwin-amd64 linux-amd64 windows-amd64
	
install: check-build doc
	install -m 755 -d         $(PREFIX)/bin/
	install -m 755 $(BIN)     $(PREFIX)/bin/
	install -m 755 -d         $(PREFIX)/share/licenses/$(NAME)/
	install -m 644 $(LICENSE) $(PREFIX)/share/licenses/$(NAME)/
	install -m 755 -d         $(PREFIX)/share/doc/$(NAME)/
	install -m 644 $(README)  $(PREFIX)/share/doc/$(NAME)/

check-install:
	test -e $(PREFIX)/bin/$(NAME) || \
	test -d $(PREFIX)/share/licenses/$(NAME) || \
	test -d $(PREFIX)/share/doc/$(NAME)

uninstall: check-install
	rm -f  $(PREFIX)/bin/$(NAME)
	rm -rf $(PREFIX)/share/licenses/$(NAME)
	rm -rf $(PREFIX)/share/doc/$(NAME)

clean:
	rm -rf $(BINDIR)
	
test: check-build
	$(TEST) $(BIN)

