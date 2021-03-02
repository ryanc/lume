LUME_VERSION ?= $(shell git describe --tags --always)
LDFLAGS := ${LDFLAGS} -X git.kill0.net/chill9/lume/cmd.Version=${LUME_VERSION}

ifeq ($(OS), Windows_NT)
    EXE=lume.exe
	RM=del /f /q
else
    EXE=lume
    EXE=rm -f
endif

.PHONY: build
build:
	go build -o ${EXE} -ldflags="${LDFLAGS}" ./cmd/lume

.PHONY: clean
clean:
	${RM} ${EXE}