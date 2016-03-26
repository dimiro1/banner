[![Build Status](https://travis-ci.org/dimiro1/banner.svg?branch=master)](https://travis-ci.org/dimiro1/banner)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimiro1/banner)](https://goreportcard.com/report/github.com/dimiro1/banner)
[![GoDoc](https://godoc.org/github.com/dimiro1/banner?status.svg)](https://godoc.org/github.com/dimiro1/banner)

# Banner

Add beautiful banners into your Go applications

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

# API

The public API is composed of two things, the method **banner.Init** and the command line flags in autoload. I do not have any plans to break compatibility, but I recommend you to vendor this dependency in your project, as it is a good practice.

# Command line flags

```sh
$ go run main.go -h
```

should output
```
Usage of main:
  -banner string
    	banner.txt file (default "banner.txt")
  -show-banner
    	print the banner? (default true)
```

# Template

You can use the following variables in the template.

| Variable                               | Value                                         |
|----------------------------------------|-----------------------------------------------|
| ```{{ .GoVersion }}```                 | ```runtime.Version()```                       |
| ```{{ .GOOS }}```                      | ```runtime.GOOS```                            |
| ```{{ .GOARCH }}```                    | ```runtime.GOARCH```                          |
| ```{{ .NumCPU }}```                    | ```runtime.NumCPU()```                        |
| ```{{ .GOPATH }}```                    | ```os.Getenv("GOPATH")```                     |
| ```{{ .GOROOT }}```                    | ```runtime.GOROOT()```                        |
| ```{{ .Compiler }}```                  | ```runtime.Compiler```                        |
| ```{{ .Env "GOPATH" }}```              | ```os.Getenv("GOPATH")```                     |
| ```{{ .Now "Monday, 2 Jan 2006" }}```  | ```time.Now().Format("Monday, 2 Jan 2006")``` |

Please see the layout of the function **.Now** in https://github.com/golang/go/blob/f06795d9b742cf3292a0f254646c23603fc6419b/src/time/format.go#L9-L41

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
