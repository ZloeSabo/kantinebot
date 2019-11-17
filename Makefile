BINARIES := kantinebot menu new_auth

binary = $(word 1, $@)

.PHONY: $(BINARIES)
$(BINARIES):
	mkdir -p ./bin
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(binary)-linux-amd64 cmd/$(binary)/$(binary).go


.PHONY: build
build: kantinebot menu new_auth

.PHONY: clean
clean:
	rm -rf ./bin