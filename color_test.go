// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package banner

import "testing"

func Test_ansiBackground_Default(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "49")
	got := ansi.Default()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Black(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "40")
	got := ansi.Black()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Red(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "41")
	got := ansi.Red()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Green(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "42")
	got := ansi.Green()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Yellow(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "43")
	got := ansi.Yellow()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Blue(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "44")
	got := ansi.Blue()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Magenta(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "45")
	got := ansi.Magenta()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Cyan(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "46")
	got := ansi.Cyan()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_White(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "47")
	got := ansi.White()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightBlack(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "100")
	got := ansi.BrightBlack()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightRed(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "101")
	got := ansi.BrightRed()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightGreen(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "102")
	got := ansi.BrightGreen()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightYellow(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "103")
	got := ansi.BrightYellow()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightBlue(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "104")
	got := ansi.BrightBlue()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightMagenta(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "105")
	got := ansi.BrightMagenta()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightCyan(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "106")
	got := ansi.BrightCyan()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightWhite(t *testing.T) {
	ansi := ansiBackground{true}
	expected := outputANSI(true, "107")
	got := ansi.BrightWhite()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Default(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "39")
	got := ansi.Default()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Black(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "30")
	got := ansi.Black()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Red(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "31")
	got := ansi.Red()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Green(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "32")
	got := ansi.Green()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Yellow(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "33")
	got := ansi.Yellow()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Blue(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "34")
	got := ansi.Blue()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Magenta(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "35")
	got := ansi.Magenta()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Cyan(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "36")
	got := ansi.Cyan()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_White(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "37")
	got := ansi.White()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightBlack(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "90")
	got := ansi.BrightBlack()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightRed(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "91")
	got := ansi.BrightRed()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightGreen(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "92")
	got := ansi.BrightGreen()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightYellow(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "93")
	got := ansi.BrightYellow()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightBlue(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "94")
	got := ansi.BrightBlue()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightMagenta(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "95")
	got := ansi.BrightMagenta()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightCyan(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "96")
	got := ansi.BrightCyan()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightWhite(t *testing.T) {
	ansi := ansiColor{true}
	expected := outputANSI(true, "97")
	got := ansi.BrightWhite()

	if got != expected {
		t.Errorf("got != expected")
	}
}
