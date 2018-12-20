# license-helper

[![GoDoc](https://godoc.org/github.com/bgaechter/license-helper?status.svg)](https://godoc.org/github.com/bgaechter/license-helper)
[![Build Status](https://travis-ci.org/bgaechter/license-helper.svg?branch=master)](https://travis-ci.org/bgaechter/license-helper)
[![Go Report Card](https://goreportcard.com/badge/github.com/bgaechter/license-helper)](https://goreportcard.com/report/github.com/bgaechter/license-helper)

`license-helper` is a little helper tool that helps you find the licenses for
your node dependencies and their dependencies. Make sure to run `npm install`
(or `yarn install` or whatever you use for pacakge management) before scanning
for dependencies.

## Installation
```
git clone github.com/bgaechter/licesne-helper
cd license-helper
go install
```

## Usage
simply run `license-helper scan [path]` to look for all package.json files
within a directory and print their license information.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/bgaechter/license-helper.

## License

The package is available as open source under the terms of the [Apache License](https://opensource.org/licenses/Apache-2.0)
