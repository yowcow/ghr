.PHONY: test

Build := _build
DarwinBuild := $(Build)/darwin-amd64
LinuxBuild := $(Build)/linux-amd64
Binary := vim-ver

all: $(DarwinBuild).tar.gz $(LinuxBuild).tar.gz

$(DarwinBuild).tar.gz: $(DarwinBuild)/$(Binary)
	tar czf $@ $<

$(LinuxBuild).tar.gz: $(LinuxBuild)/$(Binary)
	tar czf $@ $<

$(DarwinBuild)/$(Binary):
	GOOS=darwin GOARCH=amd64 go build -o $@

$(LinuxBuild)/$(Binary):
	GOOS=linux GOARCH=amd64 go build -o $@

test:
	go test

clean:
	rm $(Build)
