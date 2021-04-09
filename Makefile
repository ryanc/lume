V ?= 0
Q = $(if $(filter 1, $V),, @)
BINDIR=$(CURDIR)/bin
PREFIX=/usr
DESTDIR=bin

ifeq ($(OS), Windows_NT)
    EXE=$(BINDIR)/lume.exe
	RM=del /f /q
	BUILD_DATE=$(shell powershell Get-Date -Format "yyyy-MM-ddThh:mm:sszzz")
else
    EXE=$(BINDIR)/lume
    RM=rm -f
	BUILD_DATE=$(shell date --iso-8601=seconds)
endif

LUME_VERSION ?= $(shell git describe --tags --always)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS = \
	-X git.kill0.net/chill9/lume/cmd.Version=$(LUME_VERSION) \
	-X git.kill0.net/chill9/lume/cmd.BuildDate=$(BUILD_DATE) \
	-X git.kill0.net/chill9/lume/cmd.GitCommit=$(GIT_COMMIT)

.PHONY: build
build:
	$(Q) go build -o $(EXE) -ldflags="$(LDFLAGS)" ./cmd/lume

.PHONY: clean
clean: deb-clean
	$(Q) $(RM) $(EXE)

.PHONY: install
install:
	$(Q) install -p -D -m 0755 $(EXE) $(DESTDIR)${PREFIX}/bin/lume
	$(Q) install -p -D -m 0644 .lumerc.sample $(DESTDIR)${PREFIX}/share/lume/lumerc

DEBDIR=$(CURDIR)/debian
TMPLDIR=$(CURDIR)/packaging/debian
DEBDATE=$(shell date -R)

.PHONY: deb
deb:
	$(Q) mkdir -p $(DEBDIR)
	$(Q) sed -e 's/__VERSION__/$(LUME_VERSION)/g' $(TMPLDIR)/rules > $(DEBDIR)/rules
	$(Q) sed -e 's/__VERSION__/$(LUME_VERSION)/g' -e 's/__DATE__/$(DEBDATE)/g' $(TMPLDIR)/changelog > $(DEBDIR)/changelog
	$(Q) echo 9 > $(DEBDIR)/compat
	$(Q) cp $(TMPLDIR)/control $(DEBDIR)/control
	$(Q) dpkg-buildpackage -us -uc -b

deb-clean:
	$(Q) rm -rf $(CURDIR)/debian
