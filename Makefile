GO_SRCS = $(shell find . -type f \( -name '*.go' -and -not -iwholename '*testdata' \) )
PKG_NAME = github.com/zchee/$(notdir ${CURDIR})
CMD_NAME = ti

GO_BUILD_FLAGS ?= -v -x

TI_DEBUG ?= 

ifneq (,$(findstring debug,$(TI_DEBUG)))
GO_BUILD_FLAGS += -tags=debug
endif
ifneq (,$(findstring race,$(TI_DEBUG)))
GO_BUILD_FLAGS += -race
endif

PROF_MODE ?=

build: bin/ti

bin:
	mkdir -p ./$@

bin/ti: $(GO_SRCS) bin
	go build $(GO_BUILD_FLAGS) -o $@ $(PKG_NAME)/cmd/$(CMD_NAME)

install: $(GO_SRCS)
	go install $(GO_BUILD_FLAGS) $(PKG_NAME)/cmd/$(CMD_NAME)

uninstall:
	rm -f $(shell which $(CMD_NAME))

test/run: build
	time ./bin/ti package ..

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
