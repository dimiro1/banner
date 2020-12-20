// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/mattn/go-colorable"

	"github.com/dimiro1/banner"
)

func init() {
	templ := `{{ .Title "Banner" "" 4 }}
   {{ .AnsiColor.BrightCyan }}The title will be ascii and indented 4 spaces{{ .AnsiColor.Default }}
   GoVersion: {{ .GoVersion }}
   GOOS: {{ .GOOS }}
   GOARCH: {{ .GOARCH }}
   NumCPU: {{ .NumCPU }}
   GOPATH: {{ .GOPATH }}
   GOROOT: {{ .GOROOT }}
   Compiler: {{ .Compiler }}
   ENV: {{ .Env "GOPATH" }}
   Now: {{ .Now "Monday, 2 Jan 2006" }}
   {{ .AnsiColor.BrightGreen }}This text will appear in Green
   {{ .AnsiColor.BrightRed }}This text will appear in Red{{ .AnsiColor.Default }}`

	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
}
