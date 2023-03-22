package main

import (
	"fmt"
	"strings"

	"github.com/yaskinny/md-table-gen/internal/md"
	"github.com/yaskinny/md-table-gen/internal/values"
)

func (c *cf) Run() error {
	output := strings.Builder{}
	output.WriteString(fmt.Sprintf("## %v\n", c.ValuesHeaderName))
	for _, path := range c.ValuesFiles {
		o, err := values.RenderValueFile(path)
		if err != nil {
			return err
		}
		output.WriteString(o)
	}
	return md.WriteNewMD(output.String(), c.MDFile, c.ValuesHeaderName)
}
