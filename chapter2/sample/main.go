package main

import (
	"GoInActionCode/chapter2/sample/search"
	"log"
	"os"

	_ "GoInActionCode/chapter2/sample/matchers"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
