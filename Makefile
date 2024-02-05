default: bin/notable-client bin/notable-server

bin/notable-client: | bin
bin/notable-server: build/ui | bin
	go generate -v ./... && \
	go build -v -o ./bin ./...

.PHONY: build/ui
build/ui:
	yarn workspace @notable/ui build

bin:
	mkdir -p $@

