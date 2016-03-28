// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package banner

import (
	"os"

	"github.com/mattn/go-isatty"
)

const (
	encodeStart = "\033["
	encodeEnd   = "m"
	encodeReset = "0;39"
)

type ansiBackground struct {
	isColorEnabled bool
}

func (a ansiBackground) Default() string {
	return outputANSI(a.isColorEnabled, "49")
}

func (a ansiBackground) Black() string {
	return outputANSI(a.isColorEnabled, "40")
}

func (a ansiBackground) Red() string {
	return outputANSI(a.isColorEnabled, "41")
}

func (a ansiBackground) Green() string {
	return outputANSI(a.isColorEnabled, "42")
}

func (a ansiBackground) Yellow() string {
	return outputANSI(a.isColorEnabled, "43")
}

func (a ansiBackground) Blue() string {
	return outputANSI(a.isColorEnabled, "44")
}

func (a ansiBackground) Magenta() string {
	return outputANSI(a.isColorEnabled, "45")
}

func (a ansiBackground) Cyan() string {
	return outputANSI(a.isColorEnabled, "46")
}

func (a ansiBackground) White() string {
	return outputANSI(a.isColorEnabled, "47")
}

func (a ansiBackground) BrightBlack() string {
	return outputANSI(a.isColorEnabled, "100")
}

func (a ansiBackground) BrightRed() string {
	return outputANSI(a.isColorEnabled, "101")
}

func (a ansiBackground) BrightGreen() string {
	return outputANSI(a.isColorEnabled, "102")
}

func (a ansiBackground) BrightYellow() string {
	return outputANSI(a.isColorEnabled, "103")
}

func (a ansiBackground) BrightBlue() string {
	return outputANSI(a.isColorEnabled, "104")
}

func (a ansiBackground) BrightMagenta() string {
	return outputANSI(a.isColorEnabled, "105")
}

func (a ansiBackground) BrightCyan() string {
	return outputANSI(a.isColorEnabled, "106")
}

func (a ansiBackground) BrightWhite() string {
	return outputANSI(a.isColorEnabled, "107")
}

type ansiColor struct {
	isColorEnabled bool
}

func (a ansiColor) Default() string {
	return outputANSI(a.isColorEnabled, "39")
}

func (a ansiColor) Black() string {
	return outputANSI(a.isColorEnabled, "30")
}

func (a ansiColor) Red() string {
	return outputANSI(a.isColorEnabled, "31")
}

func (a ansiColor) Green() string {
	return outputANSI(a.isColorEnabled, "32")
}

func (a ansiColor) Yellow() string {
	return outputANSI(a.isColorEnabled, "33")
}

func (a ansiColor) Blue() string {
	return outputANSI(a.isColorEnabled, "34")
}

func (a ansiColor) Magenta() string {
	return outputANSI(a.isColorEnabled, "35")
}

func (a ansiColor) Cyan() string {
	return outputANSI(a.isColorEnabled, "36")
}

func (a ansiColor) White() string {
	return outputANSI(a.isColorEnabled, "37")
}

func (a ansiColor) BrightBlack() string {
	return outputANSI(a.isColorEnabled, "90")
}

func (a ansiColor) BrightRed() string {
	return outputANSI(a.isColorEnabled, "91")
}

func (a ansiColor) BrightGreen() string {
	return outputANSI(a.isColorEnabled, "92")
}

func (a ansiColor) BrightYellow() string {
	return outputANSI(a.isColorEnabled, "93")
}

func (a ansiColor) BrightBlue() string {
	return outputANSI(a.isColorEnabled, "94")
}

func (a ansiColor) BrightMagenta() string {
	return outputANSI(a.isColorEnabled, "95")
}

func (a ansiColor) BrightCyan() string {
	return outputANSI(a.isColorEnabled, "96")
}

func (a ansiColor) BrightWhite() string {
	return outputANSI(a.isColorEnabled, "97")
}

func (a ansiBackground) Reset() string {
	return outputANSI(a.isColorEnabled, encodeReset)
}

func outputANSI(isColorEnabled bool, code string) string {
	if isColorEnabled && isANSIEnabled() {
		return encodeStart + code + encodeEnd
	}

	return ""

}

func isANSIEnabled() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}
