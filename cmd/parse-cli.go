package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type cf struct {
	ValuesFiles      []string `short:"f" long:"value-file" description:"Value file to generate md file from(can be used multiple times)." required:"true"`
	MDFile           string   `short:"r" long:"md-file" description:"Markdown file to write genrated output to." required:"true"`
	ValuesHeaderName string   `short:"n" long:"header-name" default:"Values" description:"Header name to set for values section when"`
}

// parse flags and validate files
func parseFlagsAndValidate(args []string) (cf, error) {
	var CF cf
	if err := parseFlags(&CF, args); err != nil {
		return CF, err
	}
	if err := validateFiles(CF); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return CF, err
	}
	return CF, nil
}

// validate all the given paths to command are files and exist
func validateFiles(cliFlags cf) error {
	files := append(cliFlags.ValuesFiles, cliFlags.MDFile)
	for _, f := range files {
		fi, err := os.Stat(f)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return fmt.Errorf("expected file but %v is a directory", f)
		}
	}
	return nil
}

// parse flags
func parseFlags(cliFlags *cf, args []string) error {
	if _, err := flags.ParseArgs(cliFlags, args); err != nil {
		return err
	}
	return nil
}
