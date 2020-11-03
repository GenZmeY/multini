# WARNING: offline build not supported :(

%undefine _missing_build_ids_terminate_build
%global debug_package %{nil}

Name:     multini
Version:  0.3.0
Release:  1%{dist}
Summary:  A utility for manipulating ini files with duplicate keys
License:  MIT
Url:      https://github.com/GenZmeY/multini

Source0:  %{name}-%{version}.tar.gz

BuildRequires: golang-bin >= 1.13

Provides: %{name}

%description
A utility for easily manipulating ini files from the command line and shell scripts.

%prep
%setup -q -c

%build
make -j $(nproc) VERSION=%{version}

%install
rm -rf $RPM_BUILD_ROOT
make install PREFIX=%{buildroot}/usr

%check
make test

%clean
rm -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root)
%attr(755,root,root) %{_bindir}/%{name}
%attr(755,root,root) %dir %{_docdir}/%{name}
%attr(755,root,root) %dir %{_datadir}/licenses/%{name}
%attr(644,root,root) %{_docdir}/%{name}/README
%attr(644,root,root) %{_datadir}/licenses/%{name}/LICENSE

%changelog
* Thu Apr 30 2020 GenZmeY <genzmey@gmail.com> - 0.3.0-1
- multini-0.3.0:
- add C-style comments support.

* Thu Apr 30 2020 GenZmeY <genzmey@gmail.com> - 0.2.3-1
- multini-0.2.3:
- fixed file write for the '--inplace' option.

* Wed Apr 29 2020 GenZmeY <genzmey@gmail.com> - 0.2.2-1
- multini-0.2.2.

* Wed Apr 29 2020 GenZmeY <genzmey@gmail.com> - 0.2.1-1
- multini-0.2.1:
- fix "rename invalid cross-device link".

* Mon Apr 27 2020 GenZmeY <genzmey@gmail.com> - 0.2-1
- multini-0.2:
- fix inplace arg;
- follow symlinks by default;
- reverse flag.

* Sun Apr 26 2020 GenZmeY <genzmey@gmail.com> - 0.1-1
- First version of spec.
