SHELL := /bin/bash
TARGETS := jsoninf jsongs
PKGNAME := jsoninf

.PHONY: all
all: $(TARGETS)

%: cmd/%/main.go
	go build -o $@ $<

.PHONY: clean
clean:
	rm -f $(TARGETS)
	rm -f $(PKGNAME)_*.deb
	rm -f $(PKGNAME)*.rpm
	rm -rf packaging/debian/$(PKGNAME)/usr

.PHONY: deb
deb: all
	mkdir -p packaging/debian/$(PKGNAME)/usr/local/bin
	cp $(TARGETS) packaging/debian/$(PKGNAME)/usr/local/bin
	cd packaging/debian && fakeroot dpkg-deb --build $(PKGNAME) .
	mv packaging/debian/$(PKGNAME)_*.deb .


