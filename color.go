// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package banner

import "runtime"

const (
	encodeStart = "\033["
	encodeEnd   = "m"
	encodeReset = "0;39"
)

type ansiBackground struct{}

func (a ansiBackground) Default() string {
	return outputANSI("49")
}

func (a ansiBackground) Black() string {
	return outputANSI("40")
}

func (a ansiBackground) Red() string {
	return outputANSI("41")
}

func (a ansiBackground) Green() string {
	return outputANSI("42")
}

func (a ansiBackground) Yellow() string {
	return outputANSI("43")
}

func (a ansiBackground) Blue() string {
	return outputANSI("44")
}

func (a ansiBackground) Magenta() string {
	return outputANSI("45")
}

func (a ansiBackground) Cyan() string {
	return outputANSI("46")
}

func (a ansiBackground) White() string {
	return outputANSI("47")
}

func (a ansiBackground) BrightBlack() string {
	return outputANSI("100")
}

func (a ansiBackground) BrightRed() string {
	return outputANSI("101")
}

func (a ansiBackground) BrightGreen() string {
	return outputANSI("102")
}

func (a ansiBackground) BrightYellow() string {
	return outputANSI("103")
}

func (a ansiBackground) BrightBlue() string {
	return outputANSI("104")
}

func (a ansiBackground) BrightMagenta() string {
	return outputANSI("105")
}

func (a ansiBackground) BrightCyan() string {
	return outputANSI("106")
}

func (a ansiBackground) BrightWhite() string {
	return outputANSI("107")
}

type ansiColor struct{}

func (a ansiColor) Default() string {
	return outputANSI("39")
}

func (a ansiColor) Black() string {
	return outputANSI("30")
}

func (a ansiColor) Red() string {
	return outputANSI("31")
}

func (a ansiColor) Green() string {
	return outputANSI("32")
}

func (a ansiColor) Yellow() string {
	return outputANSI("33")
}

func (a ansiColor) Blue() string {
	return outputANSI("34")
}

func (a ansiColor) Magenta() string {
	return outputANSI("35")
}

func (a ansiColor) Cyan() string {
	return outputANSI("36")
}

func (a ansiColor) White() string {
	return outputANSI("37")
}

func (a ansiColor) BrightBlack() string {
	return outputANSI("90")
}

func (a ansiColor) BrightRed() string {
	return outputANSI("91")
}

func (a ansiColor) BrightGreen() string {
	return outputANSI("92")
}

func (a ansiColor) BrightYellow() string {
	return outputANSI("93")
}

func (a ansiColor) BrightBlue() string {
	return outputANSI("94")
}

func (a ansiColor) BrightMagenta() string {
	return outputANSI("95")
}

func (a ansiColor) BrightCyan() string {
	return outputANSI("96")
}

func (a ansiColor) BrightWhite() string {
	return outputANSI("97")
}

func outputANSI(code string) string {
	if isANSIEnabled() {
		return encodeStart + code + encodeEnd
	}

	return ""
}

func isANSIEnabled() bool {
	// Windows is always disabled.
	return isTerminal() && (runtime.GOOS != "windows")
}
