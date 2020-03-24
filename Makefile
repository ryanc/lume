.PHONY: build
build:
	go build -o lume ./cmd/lume

.PHONY: clean
clean:
	rm -f ./lifx
