REV = $(shell git log -n 1| head -n1 | awk '{print substr($$2, 0, 7)}')

BUILD = _build
BUILD_LINUX = $(BUILD)/linux-amd64
BUILD_DARWIN = $(BUILD)/darwin-amd64

all:
	$(MAKE) -j 4 \
		ghr \
		_build/linux-amd64/ghr-$(REV)-linux-amd64.tar.gz \
		_build/darwin-amd64/ghr-$(REV)-darwin-amd64.tar.gz

ghr:
	go build -ldflags '-X main.Version=$(REV)'

test:
	go test ./...

clean:
	rm -rf ghr $(BUILD)

$(BUILD_LINUX)/ghr-%-linux-amd64.tar.gz: $(BUILD_LINUX)/ghr-%-linux-amd64
	cd $(dir $<) && \
		tar czf $(notdir $@) $(notdir $<)

$(BUILD_LINUX)/ghr-%-linux-amd64: $(BUILD_LINUX)
	mkdir -p $@
	GOOS=linux GOARCH=amd64 go build -ldflags '-X main.Version=$*' -o $@/ghr

$(BUILD_DARWIN)/ghr-%-darwin-amd64.tar.gz: $(BUILD_DARWIN)/ghr-%-darwin-amd64
	cd $(dir $<) && \
		tar czf $(notdir $@) $(notdir $<)

$(BUILD_DARWIN)/ghr-%-darwin-amd64: $(BUILD_DARWIN)
	mkdir -p $@
	GOOS=darwin GOARCH=amd64 go build -ldflags '-X main.Version=$*' -o $@/ghr

$(BUILD_LINUX): $(BUILD)
	mkdir -p $@

$(BUILD_DARWIN): $(BUILD)
	mkdir -p $@

$(BUILD):
	mkdir -p $@

.PHONY: all test clean
