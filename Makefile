.PHONY: build
build:
	go build -o lume ./cmd

.PHONY: clean
clean:
	rm -f ./lifx
