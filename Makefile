.PHONY: build clean

GO_FILES := $(shell find . -name "*.go")
GO_MOD := go.mod

build: docstartsym

docstartsym: $(GO_FILES) $(GO_MOD)
	go build -o docstartsym ./cmd/docstartsym

clean:
	rm -f docstartsym 