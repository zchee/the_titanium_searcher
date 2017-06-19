GO_SRCS = $(shell find . -type f \( -name '*.go' -and -not -iwholename '*testdata' \) )
PKG_NAME = github.com/zchee/$(notdir ${CURDIR})
CMD_NAME = ti

GO_BUILD_FLAGS ?= -v -x

ifneq ($(TI_DEBUG),)
GO_BUILD_FLAGS += -tags=debug
endif

build: bin/ti

bin:
	mkdir -p ./$@

bin/ti: $(GO_SRCS) bin
	go build $(GO_BUILD_FLAGS) -o $@ $(PKG_NAME)/cmd/$(CMD_NAME)

install: $(GO_SRCS)
	go install $(GO_BUILD_FLAGS) $(PKG_NAME)/cmd/$(CMD_NAME)

uninstall:
	rm -f $(shell which $(CMD_NAME))

clean:
	rm -rf ./bin

.PHONY: build install uninstall clean
