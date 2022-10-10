# Go FeatWS

Forked from [Ini](go-ini/ini)

[![GitHub Workflow Status](https://img.shields.io/github/checks-status/bancodobrasil/go-featws/main?logo=github&style=for-the-badge)](https://github.com/bancodobrasil/go-featws/actions?query=branch%3Amain)
[![codecov](https://img.shields.io/codecov/c/github/bancodobrasil/go-featws/master?logo=codecov&style=for-the-badge)](https://codecov.io/gh/bancodobrasil/go-featws)
[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/bancodobrasil/go-featws?tab=doc)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?style=for-the-badge&logo=sourcegraph)](https://sourcegraph.com/github.com/bancodobrasil/go-featws)

<!--![](https://avatars0.githubusercontent.com/u/10216035?v=3&s=200) -->

Package go-featws provides FeatWS file read and write functionality in Go.

## Features

- Load from multiple data sources(file, `[]byte`, `io.Reader` and `io.ReadCloser`) with overwrites.
- Read with recursion values.
- Read with parent-child sections.
- Read with auto-increment key names.
- Read with multiple-line values.
- Read with tons of helper methods.
- Read and convert values to Go types.
- Read and **WRITE** comments of sections and keys.
- Manipulate sections, keys and comments with ease.
- Keep sections and keys in order as you parse and save.

## Installation

The minimum requirement of Go is **1.13**.

```sh
$ go get github.com/bancodobrasil/go-featws
```

Please add `-u` flag to update in the future.

## Getting Help

- [Getting Started](https://github.com/bancodobrasil/featws-transpiler)

## License

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.
