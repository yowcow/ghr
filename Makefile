BUILD = _build

OS = linux darwin
BIN = vimver $(foreach os,$(OS),$(BUILD)/$(os)/vimver)
ARCHIVE = $(foreach os,$(OS),$(BUILD)/vimver-$(os).tar.gz)

all:
	$(MAKE) -j 4 $(BIN)

$(BUILD)/vimver-%.tar.gz: $(BUILD)/%/vimver
	cd $(dir $<) && \
		tar czf $(notdir $@) $(notdir $<)

$(BUILD)/%/vimver:
	mkdir -p $(dir $@)
	GOOS=$* GOARCH=amd64 go build -ldflags '-X main.Version=$(VERSION)' -o $@

vimver:
	go build -ldflags '-X main.Version=$(VERSION)'

archive:
	$(MAKE) -j 4 $(ARCHIVE)

test:
	go test ./...

clean:
	rm -rf $(BIN) $(ARCHIVE)

.PHONY: all archive test clean
