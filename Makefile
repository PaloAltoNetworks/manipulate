include apomock.mk

test: apotest
clean: apoclean_vendor apoclean_apomock

build:
	go build
