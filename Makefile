BUILD = _build

OS = linux darwin
BIN = ghr $(foreach os,$(OS),$(BUILD)/$(os)/ghr)
ARCHIVE = $(foreach os,$(OS),$(BUILD)/ghr-$(os).tar.gz)

all:
	$(MAKE) -j 4 $(BIN)

$(BUILD)/ghr-%.tar.gz: $(BUILD)/%/ghr
	cd $(dir $<) && \
		tar czf $(notdir $@) $(notdir $<)

$(BUILD)/%/ghr:
	mkdir -p $(dir $@)
	GOOS=$* GOARCH=amd64 go build -ldflags '-X main.Version=$(VERSION)' -o $@

ghr:
	go build -ldflags '-X main.Version=$(VERSION)'

archive:
	$(MAKE) -j 4 $(ARCHIVE)

test:
	go test ./...

clean:
	rm -rf $(BIN) $(ARCHIVE)

.PHONY: all archive test clean
