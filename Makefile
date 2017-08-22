.PHONY: test

Bin := vim-ver

all: $(Bin)

$(Bin):
	go build -o $@

test:
	go test

clean:
	rm $(Bin)
