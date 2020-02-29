.PHONY: build
build:
	go build -o lifx ./cmd

.PHONY: clean
clean:
	rm -f ./lifx
