GO := /home/k/go/bin/go1.11.4

.PHONY : clean build

all : build

build:
	$(GO) mod tidy
	$(GO) build

clean: dedupPATH
	rm dedupPATH
