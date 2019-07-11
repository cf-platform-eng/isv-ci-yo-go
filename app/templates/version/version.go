package version

import (
	"fmt"

	"<%= repoURL %>"
)

type VersionOpt struct {
}

var Version = "dev"

func (_ *VersionOpt) Execute(args []string) error {
	fmt.Printf("%s version: %s\n", <%= appName %>.APP_NAME, Version)
	return nil
}
