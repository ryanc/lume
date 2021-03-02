ifeq ($(OS), Windows_NT)
    EXE=lume.exe
	RM=del /f
else
    EXE=lume
    EXE=rm -f
endif

.PHONY: build
build:
	go build -o ${EXE} ./cmd/lume

.PHONY: clean
clean:
	${RM} ${EXE}
