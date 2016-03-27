// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package banner

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"
	"time"
)

func Test_printBanner(t *testing.T) {
	bannerContent := `Hello, {{ .GoVersion }}, {{ .GOOS }}, {{ .GOARCH }}, {{ .NumCPU }}, {{ .GOPATH }}, {{ .GOROOT }}, {{ .Compiler }}`
	var buffer bytes.Buffer

	version := runtime.Version()
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	numCPU := runtime.NumCPU()
	gopath := os.Getenv("GOPATH")
	goroot := runtime.GOROOT()
	compiler := runtime.Compiler

	expected := fmt.Sprintf(`Hello, %s, %s, %s, %d, %s, %s, %s`, version, goos, goarch, numCPU, gopath, goroot, compiler)

	show(&buffer, true, bannerContent)

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_invalid(t *testing.T) {
	var buffer bytes.Buffer

	expected := ""

	show(&buffer, true, "{{}")

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_flags(t *testing.T) {
	var buffer bytes.Buffer

	Init(&buffer, true, true, bytes.NewBufferString("Test Banner"))

	expected := "Test Banner"

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_invalid_reader(t *testing.T) {
	var buffer bytes.Buffer

	Init(&buffer, true, true, nil)

	expected := ""

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_closed_reader(t *testing.T) {
	var buffer bytes.Buffer

	// Testing an error while reading a closed file.
	in, err := os.Open("test-banner.txt")

	if err != nil {
		t.Error(err)
	}

	in.Close()

	Init(&buffer, true, true, in)

	expected := ""

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_printBanner_banner_disabled(t *testing.T) {
	var buffer bytes.Buffer

	Init(&buffer, false, true, bytes.NewBufferString("Test Banner"))

	expected := ""

	result, err := ioutil.ReadAll(&buffer)

	if err != nil {
		t.Error(err)
	}

	if string(result) != expected {
		t.Errorf("result != expected, got %s", string(result))
	}
}

func Test_vars_Env(t *testing.T) {
	v := vars{}
	gopath := v.Env("GOPATH")

	expected := os.Getenv("GOPATH")

	if gopath != expected {
		t.Errorf("gopath != expected, got %s", gopath)
	}
}

func Test_vars_Now(t *testing.T) {
	v := vars{}
	gopath := v.Now("Monday, 2 Jan 2006")

	expected := time.Now().Format("Monday, 2 Jan 2006")

	if gopath != expected {
		t.Errorf("gopath != expected, got %s", gopath)
	}
}

func Test_SetLog(t *testing.T) {
	oldLogger := logger
	SetLog(log.New(os.Stderr, "", log.LstdFlags))

	if oldLogger == logger {
		t.Errorf("logger was changed, must not be equal")
	}
}

func Test_SetLog_nil(t *testing.T) {
	oldLogger := logger
	SetLog(nil)

	if oldLogger != logger {
		t.Errorf("logger was not changed, must be equal")
	}
}
