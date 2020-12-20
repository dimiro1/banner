// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package autoload configure the banner loader with defaults
// Import the package. Thats it.
package autoload

import (
	"flag"
	"os"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

func init() {
	var (
		filename       string
		isEnabled      bool
		isColorEnabled bool
	)

	flag.StringVar(&filename, "banner", "", "read from a txt file")
	flag.BoolVar(&isEnabled, "show-banner", true, "print the banner?")
	flag.BoolVar(&isColorEnabled, "ansi", true, "ansi colors enabled?")

	flag.Parse()

	if filename != "" {
		in, err := os.Open(filename)

		if in != nil {
			defer in.Close()
		}

		if err != nil {
			return
		}

		banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, in)
		return
	}

}
