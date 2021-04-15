V ?= 0
Q = $(if $(filter 1, $V),, @)
BINDIR=$(CURDIR)/bin
PREFIX=/usr
DESTDIR=bin
BUILDDIR=$(CURDIR)/build
MANDIR=$(PREFIX)/share/man/man1

DEBBUILDDIR=$(BUILDDIR)/deb
DEBTMPLDIR=$(CURDIR)/packaging/debian
DEBDATE=$(shell date -R)
DEBORIGSRC=lume_$(DEBVERSION).orig.tar.xz
DEBORIGSRCDIR=lume-$(DEBVERSION)

RPMVERSION=$(subst -,_,$(LUME_VERSION))
RPMBUILDDIR=$(BUILDDIR)/rpm
RPMTMPLDIR=$(CURDIR)/packaging/rpm
RPMDATE=$(shell date "+%a %b %d %Y")
RPMORIGSRC=lume-$(RPMVERSION).tar.xz
RPMORIGSRCDIR=lume-$(RPMVERSION)

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
clean: deb-clean rpm-clean
	$(Q) $(RM) $(EXE)

install-man:
	install -p -D -m 0644 lume.1 $(DESTDIR)$(MANDIR)/lume.1

.PHONY: install
install: install-man
	$(Q) install -p -D -m 0755 $(EXE) $(DESTDIR)${PREFIX}/bin/lume
	$(Q) install -p -D -m 0644 .lumerc.sample $(DESTDIR)${PREFIX}/share/lume/lumerc

.PHONY: deb
deb:
	$(Q) mkdir -p $(DEBBUILDDIR)
	$(Q) git archive --format tar --prefix lume-$(DEBVERSION)/ $(LUME_VERSION) | xz > $(DEBBUILDDIR)/$(DEBORIGSRC)
	$(Q) tar xf $(DEBBUILDDIR)/$(DEBORIGSRC) -C $(DEBBUILDDIR)
	$(Q) mkdir $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian
	$(Q) mkdir $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/source
	$(Q) sed -e 's/__VERSION__/$(DEBVERSION)/g' $(DEBTMPLDIR)/rules > $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/rules
	$(Q) chmod 0755 $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/rules
	$(Q) sed -e 's/__VERSION__/$(DEBVERSION)/g' -e 's/__DATE__/$(DEBDATE)/g' $(DEBTMPLDIR)/changelog > $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/changelog
	$(Q) echo 10 > $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/compat
	$(Q) echo "3.0 (quilt)" > $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/source/format
	$(Q) cp $(DEBTMPLDIR)/control $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/control
	$(Q) cp $(DEBTMPLDIR)/copyright $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/copyright
	$(Q) cp $(DEBTMPLDIR)/lume.manpages $(DEBBUILDDIR)/$(DEBORIGSRCDIR)/debian/lume.manpages
	$(Q) cd $(DEBBUILDDIR)/$(DEBORIGSRCDIR) && dpkg-buildpackage -us -uc
	$(Q) mv $(DEBBUILDDIR)/*.dsc $(BUILDDIR)
	$(Q) mv $(DEBBUILDDIR)/*.changes $(BUILDDIR)
	$(Q) mv $(DEBBUILDDIR)/*.buildinfo $(BUILDDIR)
	$(Q) mv $(DEBBUILDDIR)/*.deb $(BUILDDIR)
	$(Q) mv $(DEBBUILDDIR)/*.tar.* $(BUILDDIR)

.PHONY: rpm
rpm:
	$(Q) mkdir -p $(RPMBUILDDIR)/SPECS
	$(Q) mkdir -p $(RPMBUILDDIR)/SOURCES
	$(Q) sed -e 's/__VERSION__/$(RPMVERSION)/g' -e 's/__DATE__/$(RPMDATE)/g' $(RPMTMPLDIR)/lume.spec > $(RPMBUILDDIR)/SPECS/lume.spec
	$(Q) git archive --format tar --prefix $(RPMORIGSRCDIR)/ $(LUME_VERSION) | xz > $(RPMBUILDDIR)/SOURCES/$(RPMORIGSRC)
	$(Q) rpmbuild --define "_topdir $(RPMBUILDDIR)"  -ba $(RPMBUILDDIR)/SPECS/lume.spec
	$(Q) mv $(RPMBUILDDIR)/RPMS/*/*.rpm $(BUILDDIR)
	$(Q) mv $(RPMBUILDDIR)/SRPMS/*.rpm $(BUILDDIR)

deb-clean:
	$(Q) rm -rf $(DEBBUILDDIR)
	$(Q) rm -f $(BUILDDIR)/*.dsc
	$(Q) rm -f $(BUILDDIR)/*.changes
	$(Q) rm -f $(BUILDDIR)/*.buildinfo
	$(Q) rm -f $(BUILDDIR)/*.deb
	$(Q) rm -f $(BUILDDIR)/*.tar.*

rpm-clean:
	$(Q) rm -rf $(RPMBUILDDIR)
	$(Q) rm -f $(BUILDDIR)/*.rpm
