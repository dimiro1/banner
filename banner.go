// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package banner

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/common-nighthawk/go-figure"
)

var logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

// SetLog permits the user change the default logger instance
func SetLog(l *log.Logger) {
	if l == nil {
		return
	}
	logger = l
}

type vars struct {
	GoVersion      string
	GOOS           string
	GOARCH         string
	NumCPU         int
	GOPATH         string
	GOROOT         string
	Compiler       string
	AnsiColor      ansiColor
	AnsiBackground ansiBackground
}

func (v vars) Env(env string) string {
	return os.Getenv(env)
}

// See https://github.com/golang/go/blob/f06795d9b742cf3292a0f254646c23603fc6419b/src/time/format.go#L9-L41
func (v vars) Now(layout string) string {
	return time.Now().Format(layout)
}

// Create ASCII art from title
func (v vars) Title(title, font string, indentW int) string {
	// TODO(phanirithvij) variadic args
	fig := figure.NewFigure(title, font, true)
	text := fig.String()
	// TODO(phanirithvij) Use ColorString as soon as PR
	// https://github.com/common-nighthawk/go-figure/pull/16 lands
	// if color != "" {
	// 	fig = figure.NewColorFigure(title, font, color, true)
	// }
	// text := fig.ColorString()
	// TODO(phanirithvij) figure out indetation https://stackoverflow.com/q/46066524/8608146
	return indent(indentW, text)
}

// See https://github.com/Masterminds/sprig/blob/4241ae82dd07e5d3391a83c25b79c04450eb22b0/strings.go#L109
func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

// InitString load the banner from a string and prints it to output
// All errors are ignored, the application will not print the banner in case of error.
func InitString(out io.Writer, isEnabled, isColorEnabled bool, templ string) {
	Init(out, isEnabled, isColorEnabled, bytes.NewBufferString(templ))
}

// Init load the banner and prints it to output
// All errors are ignored, the application will not print the banner in case of error.
func Init(out io.Writer, isEnabled, isColorEnabled bool, in io.Reader) {
	if !isEnabled {
		logger.Println("The banner is not enabled.")
		return
	}

	if in == nil {
		logger.Println("The input is nil")
		return
	}

	banner, err := ioutil.ReadAll(in)

	if err != nil {
		logger.Printf("Error trying to read the banner, err: %v", err)
		return
	}

	show(out, isColorEnabled, string(banner))
}

func show(out io.Writer, isColorEnabled bool, content string) {
	t, err := template.New("banner").Parse(content)

	if err != nil {
		logger.Printf("Error trying to parse the banner file, err: %v", err)
		return
	}

	t.Execute(out, vars{
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		os.Getenv("GOPATH"),
		runtime.GOROOT(),
		runtime.Compiler,
		ansiColor{isColorEnabled},
		ansiBackground{isColorEnabled},
	})
}
