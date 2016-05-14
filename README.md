# Go language bindings for ADSI

This package provides access to the Active Directory Service Interfaces that are
available through the Windows component object model API. This package should
compile on any platform but implementations are only provided for Windows.

The `adsi` package provides high level and idiomatic access to ADSI. It in turn
relies on the `api` package, which handles the low level details of COM binding
and is analogous to Go's `syscall` package.

This project is a work in progress. Only a small subset of the available
interfaces have been implemented.
