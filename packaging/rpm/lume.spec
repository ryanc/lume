Name:           lume
Version:        __VERSION__
Release:        1%{?dist}
Summary:        A CLI tool for the LIFX HTTP API

License:        MPL
URL:            https://git.kill0.net/chill9/lume
Source:         %{name}-%{version}.tar.xz

%global debug_package %{nil}

%description

%prep
%setup

%build
%make_build

%install
%make_install DESTDIR=%{buildroot}


%files
%{_bindir}/lume
%license LICENSE
/usr/share/lume/lumerc


%changelog
* __DATE__ Ryan Cavicchioni <ryan@cavi.cc>
- lume __VERSION__
