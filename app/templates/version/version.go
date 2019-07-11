package version

import (
	"fmt"
	"io"

	"<%= repoURL %>"
)


type VersionOpt struct {
	Out io.Writer
}

var Version = "dev"

func (opts *VersionOpt) Execute(args []string) error {
	_, err := fmt.Fprintf(opts.Out, "%s version: %s\n", <%= appName %>.APP_NAME, Version)
	return err
}

