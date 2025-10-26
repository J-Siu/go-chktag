# go-chktag [![Paypal donate](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/donate/?business=HZF49NM9D35SJ&no_recurring=0&currency_code=CAD)

Command line tool checking `version.go`, `CHANGELOG.md` and git tag before tagging.

- [Install](#install)
- [What It Does](#what-it-does)
- [Usage](#usage)
- [License](#license)

<!--more-->

### Install

Go install

```sh
go install github.com/J-Siu/go-chktag@latest
```

Download

- https://github.com/J-Siu/go-chktag/releases

### What It Does

> This is a very simple and opinionated cli check before applying git tag.

- Check the tag
  - does not exist in git tag
  - exists in `version.go`
  - exists in `CHANGELOG.md`

### Usage

```sh
Show changelog.md, version.go and git tag.
Use -t to specify tag version.

Usage:
  go-chktag [flags]

Flags:
  -d, --debug        Enable debug
  -h, --help         help for go-chktag
  -t, --tag string   check specific tag
  -v, --version      version for go-chktag
```

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
