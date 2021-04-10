V ?= 0
Q = $(if $(filter 1, $V),, @)
BINDIR=$(CURDIR)/bin
PREFIX=/usr
DESTDIR=bin
BUILDDIR=$(CURDIR)/build

DEBBUILDDIR=$(BUILDDIR)/deb
DEBTMPLDIR=$(CURDIR)/packaging/debian
DEBDATE=$(shell date -R)
DEBORIGSRC=lume_$(DEBVERSION).orig.tar.xz
DEBORIGSRCDIR=lume-$(DEBVERSION)

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
GIT_TAG=$(shell git describe --tags --abbrev=0)
LDFLAGS = \
	-X git.kill0.net/chill9/lume/cmd.Version=$(LUME_VERSION) \
	-X git.kill0.net/chill9/lume/cmd.BuildDate=$(BUILD_DATE) \
	-X git.kill0.net/chill9/lume/cmd.GitCommit=$(GIT_COMMIT)

ifneq (,$(findstring -,$(LUME_VERSION)))
	DEBVERSION=$(GIT_TAG)+git$(shell date +%Y%m%d)+$(GIT_COMMIT)
else
	DEBVERSION=$(LUME_VERSION)
endif

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

.PHONY: deb
deb:
	$(Q) mkdir -p $(DEBBUILDDIR)
	$(Q) git archive --format tar --prefix lume-$(DEBVERSION)/ $(LUME_VERSION) | xz > $(DEBBUILDDIR)/$(DEBORIGSRC)
	$(Q) tar xf $(DEBBUILDDIR)/$(DEBORIGSRC) -C $(DEBBUILDDIR)
	$(Q) mkdir $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian
	$(Q) sed -e 's/__VERSION__/$(DEBVERSION)/g' $(DEBTMPLDIR)/rules > $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/rules
	$(Q) sed -e 's/__VERSION__/$(DEBVERSION)/g' -e 's/__DATE__/$(DEBDATE)/g' $(DEBTMPLDIR)/changelog > $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/changelog
	$(Q) echo 9 > $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/compat
	$(Q) cp $(DEBTMPLDIR)/control $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/control
	$(Q) dpkg-source -b $(DEBBUILDDIR)/$(DEBORIGSRCDIR)
	$(Q) cd $(DEBBUILDDIR)/$(DEBORIGSRCDIR) && dpkg-buildpackage -us -uc

deb-clean:
	$(Q) rm -rf $(CURDIR)/build/deb
