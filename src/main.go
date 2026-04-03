/*
Copyright © 2025, [Pavel Milanov]: forge cli utility
*/
package main

import (
	"os"

	"github.com/PavelMilanov/forge/cmd"
)

// main точка входа в приложение forge CLI.
func main() {

	cmd.Execute()
	os.Exit(0)
}
