// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package banner

import "testing"

func Test_ansiBackground_Default(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("49")
	got := ansi.Default()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Black(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("40")
	got := ansi.Black()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Red(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("41")
	got := ansi.Red()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Green(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("42")
	got := ansi.Green()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Yellow(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("43")
	got := ansi.Yellow()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Blue(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("44")
	got := ansi.Blue()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Magenta(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("45")
	got := ansi.Magenta()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_Cyan(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("46")
	got := ansi.Cyan()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_White(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("47")
	got := ansi.White()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightBlack(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("100")
	got := ansi.BrightBlack()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightRed(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("101")
	got := ansi.BrightRed()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightGreen(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("102")
	got := ansi.BrightGreen()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightYellow(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("103")
	got := ansi.BrightYellow()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightBlue(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("104")
	got := ansi.BrightBlue()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightMagenta(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("105")
	got := ansi.BrightMagenta()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightCyan(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("106")
	got := ansi.BrightCyan()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiBackground_BrightWhite(t *testing.T) {
	ansi := ansiBackground{}
	expected := outputANSI("107")
	got := ansi.BrightWhite()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Default(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("39")
	got := ansi.Default()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Black(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("30")
	got := ansi.Black()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Red(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("31")
	got := ansi.Red()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Green(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("32")
	got := ansi.Green()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Yellow(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("33")
	got := ansi.Yellow()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Blue(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("34")
	got := ansi.Blue()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Magenta(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("35")
	got := ansi.Magenta()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_Cyan(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("36")
	got := ansi.Cyan()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_White(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("37")
	got := ansi.White()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightBlack(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("90")
	got := ansi.BrightBlack()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightRed(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("91")
	got := ansi.BrightRed()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightGreen(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("92")
	got := ansi.BrightGreen()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightYellow(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("93")
	got := ansi.BrightYellow()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightBlue(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("94")
	got := ansi.BrightBlue()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightMagenta(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("95")
	got := ansi.BrightMagenta()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightCyan(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("96")
	got := ansi.BrightCyan()

	if got != expected {
		t.Errorf("got != expected")
	}
}

func Test_ansiColor_BrightWhite(t *testing.T) {
	ansi := ansiColor{}
	expected := outputANSI("97")
	got := ansi.BrightWhite()

	if got != expected {
		t.Errorf("got != expected")
	}
}
