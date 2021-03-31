V ?= 0
Q = $(if $(filter 1, $V),, @)

BIN_DIR=./bin

ifeq ($(OS), Windows_NT)
    EXE=$(BIN_DIR)/lume.exe
	RM=del /f /q
	BUILD_DATE=$(shell powershell Get-Date -Format "yyyy-MM-ddThh:mm:sszzz")
else
    EXE=$(BIN_DIR)/lume
    RM=rm -f
	BUILD_DATE=$(shell date --iso-8601=seconds)
endif

LUME_VERSION ?= $(shell git describe --tags --always)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := $(LDFLAGS) \
	-X git.kill0.net/chill9/lume/cmd.Version=$(LUME_VERSION) \
	-X git.kill0.net/chill9/lume/cmd.BuildDate=$(BUILD_DATE) \
	-X git.kill0.net/chill9/lume/cmd.GitCommit=$(GIT_COMMIT)

.PHONY: build
build:
	$(Q) go build -o $(EXE) -ldflags="$(LDFLAGS)" ./cmd/lume

.PHONY: clean
clean:
	$(Q) $(RM) $(EXE)
