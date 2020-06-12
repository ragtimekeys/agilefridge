package main

import (
	"os"

	"github.com/ragtimekeys/agilefridge/internal/agilefridge"
)

func main() {
	os.Exit(agilefridge.Run(agilefridge.GetConfigurationFromEnvironment()))
}
