
ifeq ($(OS), Windows_NT)
    EXE=lume.exe
	RM=del /f /q
	BUILD_DATE=$(shell powershell Get-Date -Format "yyyy-MM-ddThh:mm:sszzz")
else
    EXE=lume
    EXE=rm -f
	BUILD_DATE=$(shell date --iso-8601=seconds)
endif

LUME_VERSION ?= $(shell git describe --tags --always)
LDFLAGS := ${LDFLAGS} \
	-X git.kill0.net/chill9/lume/cmd.Version=${LUME_VERSION} \
	-X git.kill0.net/chill9/lume/cmd.BuildDate=${BUILD_DATE}

.PHONY: build
build:
	go build -o ${EXE} -ldflags="${LDFLAGS}" ./cmd/lume

.PHONY: clean
clean:
	${RM} ${EXE}