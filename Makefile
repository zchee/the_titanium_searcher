GO_SRCS = $(shell find . -type f \( -name '*.go' -and -not -iwholename '*testdata' \) )
PKG_NAME = github.com/zchee/$(notdir ${CURDIR})
CMD_NAME = ti

build: bin/ti

bin:
	mkdir -p ./$@

bin/ti: $(GO_SRCS) bin
	go build -v -x -o $@ $(PKG_NAME)/cmd/$(CMD_NAME)

install: $(GO_SRCS)
	go install -v -x $(PKG_NAME)/cmd/$(CMD_NAME)

uninstall:
	rm -f $(shell which $(CMD_NAME))

clean:
	rm -rf ./bin

.PHONY: build install uninstall clean
