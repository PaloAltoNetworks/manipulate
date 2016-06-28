include apomock.mk

clean: apoclean_vendor apoclean_apomock

build:
	go build

init: apoinit
test: apotest
release:
