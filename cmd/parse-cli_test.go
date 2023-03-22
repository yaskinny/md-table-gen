package main

import (
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	td := []struct {
		args     []string
		mustFail bool
		desc     string
	}{
		{
			args: []string{
				"-f",
				"file1.yaml",
				"-r",
				"readme.md",
			},
			mustFail: false,
			desc:     "everything is okay",
		}, {
			args:     []string{},
			mustFail: true,
			desc:     "at least one value file and a target md file should be defined",
		},
		{
			args: []string{
				"-f",
				"file1.yaml",
			},
			mustFail: true,
			desc:     "target md file is not set",
		}, {
			args: []string{
				"-r",
				"file1.md",
			},
			mustFail: true,
			desc:     "at least one value file is needed",
		},
	}

	for _, d := range td {
		if err := parseFlags(&cf{}, d.args); err != nil && !d.mustFail {
			t.Errorf("test failed, err: %v, desc: %v, args: %v, mustFail: %v\n", err, d.desc, d.args, d.mustFail)
		}
	}
}

func TestValidateFiles(t *testing.T) {
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}
	defer os.Remove(dir)

	file, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	file.Close()
	defer os.Remove(file.Name())

	td := []struct {
		args     []string
		desc     string
		mustFail bool
	}{
		{
			args: []string{
				"-f",
				file.Name(),
				"-r",
				file.Name(),
			},
			desc:     "both are files, shouldn't raise any error",
			mustFail: false,
		}, {
			args: []string{
				"-f",
				dir,
				"-r",
				file.Name(),
			},
			desc:     "use directory instead of file for -f option",
			mustFail: true,
		},
	}
	for _, d := range td {
		var cliFlags cf
		parseFlags(&cliFlags, d.args)
		if err := validateFiles(cliFlags); err != nil && !d.mustFail {
			t.Errorf("test failed, err: %v, desc: %v, args: %v, mustFail: %v\n", err, d.desc, d.args, d.mustFail)
		}
	}
}
