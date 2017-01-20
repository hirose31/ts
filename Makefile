SOURCES = $(shell echo *.go)
VERSION = $(shell grep '^const Version' main.go | awk '{print $$4}')

all:
	@echo ts/v$(VERSION)

ts: $(SORUCES)
	go build -o $@ $(SOURCES)

clean:
	$(RM) *~ ts
