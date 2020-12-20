[![Build Status](https://travis-ci.org/dimiro1/banner.svg?branch=master)](https://travis-ci.org/dimiro1/banner)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimiro1/banner)](https://goreportcard.com/report/github.com/dimiro1/banner)
[![GoDoc](https://godoc.org/github.com/dimiro1/banner?status.svg)](https://godoc.org/github.com/dimiro1/banner)

Try browsing [the code on Sourcegraph](https://sourcegraph.com/github.com/dimiro1/banner)!

# Banner

Add beautiful banners into your Go applications

**Table of Contents**

- [Motivation](#motivation)
- [Usage](#usage)
- [API](#api)
- [Command line flags](#command-line-flags)
- [Template](#template)
- [Log](#log)
- [ASCII Banners](#ascii-banners)
- [LICENSE](#license)

# Motivation

I Like to add these startup banners on all my applications, I think it give personality to the application.

# Usage

Import the package. Thats it.

```go
package main

import _ "github.com/dimiro1/banner/autoload"

func main() {}
```

By default it look at the file **banner.txt** in the same directory. You can customize with the command line flags.

If you do not want to use the autoload package you can always fallback to the banner API

```go
package main

import (
	"bytes"
	"os"

	"github.com/dimiro1/banner"
)

func main() {
  isEnabled := true
  isColorEnabled := true
  banner.Init(os.Stdout, isEnabled, isColorEnabled, bytes.NewBufferString("My Custom Banner"))
}
```

If using windows, use go-colorable. This works in all-platforms.

```go
package main

import (
	"bytes"
	"os"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

func main() {
  isEnabled := true
  isColorEnabled := true
  banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString("My Custom Banner"))
}
```

# API

I recommend you to vendor this dependency in your project, as it is a good practice.

# Command line flags

```sh
$ go run main.go -h
```

should output

```
Usage of main:
  -ansi
    	ansi colors enabled? (default true)
  -banner string
    	banner.txt file (default "banner.txt")
  -show-banner
    	print the banner? (default true)
```

# Template

You can use the following variables in the template.

| Variable                                     | Value                                     |
| -------------------------------------------- | ----------------------------------------- |
| `{{ .Title "YourTitle" "fontname" indent }}` | `<Generated ASCII art>`                   |
| `{{ .GoVersion }}`                           | `runtime.Version()`                       |
| `{{ .GOOS }}`                                | `runtime.GOOS`                            |
| `{{ .GOARCH }}`                              | `runtime.GOARCH`                          |
| `{{ .NumCPU }}`                              | `runtime.NumCPU()`                        |
| `{{ .GOPATH }}`                              | `os.Getenv("GOPATH")`                     |
| `{{ .GOROOT }}`                              | `runtime.GOROOT()`                        |
| `{{ .Compiler }}`                            | `runtime.Compiler`                        |
| `{{ .Env "GOPATH" }}`                        | `os.Getenv("GOPATH")`                     |
| `{{ .Now "Monday, 2 Jan 2006" }}`            | `time.Now().Format("Monday, 2 Jan 2006")` |

Please see the layout of the function **.Now** in https://github.com/golang/go/blob/f06795d9b742cf3292a0f254646c23603fc6419b/src/time/format.go#L9-L41

## Title

Title generates ascii art for you using [go-figure](https://github.com/common-nighthawk/go-figure)
Use it if you don't provide your own ascii title.

See [examples/title](./examples/title) for an example

**Note:** Must provide zero values if not using something i.e.

```go
// .Title string   string    int
// .Title title    fontname  indentation_spaces
{{ .Title "Banner" ""        0 }}
{{ .Title "Banner" "banner2" 0 }}
{{ .Title "Banner" ""        4 }}
```

The fonts available can be seen here [go-figure#supported-fonts](https://github.com/common-nighthawk/go-figure#supported-fonts)

## Colors

There are support for ANSI colors :)

| Variable                              |
| ------------------------------------- |
| `{{ .AnsiColor.Default }}`            |
| `{{ .AnsiColor.Black }}`              |
| `{{ .AnsiColor.Red }}`                |
| `{{ .AnsiColor.Green }}`              |
| `{{ .AnsiColor.Yellow }}`             |
| `{{ .AnsiColor.Blue }}`               |
| `{{ .AnsiColor.Magenta }}`            |
| `{{ .AnsiColor.Cyan }}`               |
| `{{ .AnsiColor.White }}`              |
| `{{ .AnsiColor.BrightBlack }}`        |
| `{{ .AnsiColor.BrightRed }}`          |
| `{{ .AnsiColor.BrightGreen }}`        |
| `{{ .AnsiColor.BrightYellow }}`       |
| `{{ .AnsiColor.BrightBlue }}`         |
| `{{ .AnsiColor.BrightMagenta }}`      |
| `{{ .AnsiColor.BrightCyan }}`         |
| `{{ .AnsiColor.BrightWhite }}`        |
| `{{ .AnsiBackground.Default }}`       |
| `{{ .AnsiBackground.Black }}`         |
| `{{ .AnsiBackground.Red }}`           |
| `{{ .AnsiBackground.Green }}`         |
| `{{ .AnsiBackground.Yellow }}`        |
| `{{ .AnsiBackground.Blue }}`          |
| `{{ .AnsiBackground.Magenta }}`       |
| `{{ .AnsiBackground.Cyan }}`          |
| `{{ .AnsiBackground.White }}`         |
| `{{ .AnsiBackground.BrightBlack }}`   |
| `{{ .AnsiBackground.BrightRed }}`     |
| `{{ .AnsiBackground.BrightGreen }}`   |
| `{{ .AnsiBackground.BrightYellow }}`  |
| `{{ .AnsiBackground.BrightBlue }}`    |
| `{{ .AnsiBackground.BrightMagenta }}` |
| `{{ .AnsiBackground.BrightCyan }}`    |
| `{{ .AnsiBackground.BrightWhite }}`   |

Want to see a nyancat?

```sh
$ go run examples/file/main.go -banner examples/file/nyancat.txt
```

![NyanCat Banner](banner-nyan.png?raw=true "NyanCat Banner")

## Example

```
  ____
 |  _ \
 | |_) | __ _ _ __  _ __   ___ _ __
 |  _ < / _` | '_ \| '_ \ / _ \ '__|
 | |_) | (_| | | | | | | |  __/ |
 |____/ \__,_|_| |_|_| |_|\___|_|

GoVersion: {{ .GoVersion }}
GOOS: {{ .GOOS }}
GOARCH: {{ .GOARCH }}
NumCPU: {{ .NumCPU }}
GOPATH: {{ .GOPATH }}
GOROOT: {{ .GOROOT }}
Compiler: {{ .Compiler }}
ENV: {{ .Env "GOPATH" }}
Now: {{ .Now "Monday, 2 Jan 2006" }}
```

will output something like this

```
  ____
 |  _ \
 | |_) | __ _ _ __  _ __   ___ _ __
 |  _ < / _` | '_ \| '_ \ / _ \ '__|
 | |_) | (_| | | | | | | |  __/ |
 |____/ \__,_|_| |_|_| |_|\___|_|

GoVersion: go1.6
GOOS: darwin
GOARCH: amd64
NumCPU: 4
GOPATH: /Users/claudemiro/go
GOROOT: /usr/local/Cellar/go/1.6/libexec
Compiler: gc
ENV: /Users/claudemiro/go
Now: Friday, 26 Mar 2016
```

# Log

I am using the standard golang log, but there is a function SetLog that accepts a custom log, so you can customize the way you want.

# ASCII Banners

Access http://patorjk.com/software/taag/#p=display&f=Big&t=Banner to generate ASCII banners.

Or use `{{ .Title }}` template

# LICENSE

The MIT License (MIT)

Copyright (c) 2016 Claudemiro

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
