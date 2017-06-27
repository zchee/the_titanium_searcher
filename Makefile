GO_SRCS = $(shell find . -type f \( -name '*.go' -and -not -iwholename '*testdata' \) )
PKG_NAME = github.com/zchee/$(notdir ${CURDIR})
CMD_NAME = ti

GO_BUILD_FLAGS ?= -v -x

TI_DEBUG ?= 

ifneq (,$(findstring debug,$(TI_DEBUG)))
GO_BUILD_FLAGS += -ldflags='-X=$(PKG_NAME).debug=true'
endif
ifneq (,$(findstring race,$(TI_DEBUG)))
GO_BUILD_FLAGS += -race
endif

PROF_MODE ?=

CRESET := \x1b[0m
CBLUE := \x1b[34;01m

build: bin/ti

bin:
	mkdir -p ./$@

bin/ti: $(GO_SRCS) bin
	go build $(GO_BUILD_FLAGS) -o $@ $(PKG_NAME)/cmd/$(CMD_NAME)

install: $(GO_SRCS)
	go install $(GO_BUILD_FLAGS) $(PKG_NAME)/cmd/$(CMD_NAME)

uninstall:
	rm -f $(shell which $(CMD_NAME))

test:
	go test -v $(shell go list ./...)

test/run: build
	gtime -v ./bin/ti package .. .

lint: gofmt golint govet

gofmt:
	@echo "$(CBLUE)=>$(CRESET) gofmt -e -s -l ..." && gofmt -e -s -l $(shell find . -name '*.go' -and -not -iwholename '*vendor*')
golint:
	@echo "$(CBLUE)=>$(CRESET) golint ..." && golint -set_exit_status $(shell go list ./... | grep -v -e internal/fastwalk)
govet:
	@echo "$(CBLUE)=>$(CRESET) go vet ..." && go vet $(shell go list ./...)

prof:
	rm -f *.pprof
	./bin/ti --profile $(PROF_MODE) package ..
	go tool pprof $(PROF_FLAGS) ./bin/ti *.pprof

prof/cpu: PROF_MODE = cpu
prof/cpu: PROF_FLAGS = -top -cum
prof/cpu: prof
prof/mem: PROF_MODE = mem
prof/mem: PROF_FLAGS = -top -cum --alloc_space
prof/mem: prof

clean:
	rm -rf ./bin *.pprof

.PHONY: build install uninstall test/run clean
