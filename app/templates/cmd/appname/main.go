package main

import (
	"fmt"
	"os"

	"<%= repoURL %>"

	"github.com/jessevdk/go-flags"

	"<%= repoURL %>/version"
)

var config <%= appName %>.Config
var parser = flags.NewParser(&config, flags.Default)

func main() {
	_, err := parser.AddCommand(
		"version",
		"print version",
		fmt.Sprintf("print %s version", <%= appName %>.APP_NAME),
		&version.VersionOpt{
			Out: os.Stdout,
		})
	if err != nil {
		fmt.Println("Could not add version command")
		os.Exit(1)
	}

	_, err = parser.Parse()
	if err != nil {
		os.Exit(1)
	}
}
