include apomock.mk

clean: apoclean_vendor apoclean_apomock

build:
	go build

# ci lifecycle
init: apoinit

test: apotest

release:
