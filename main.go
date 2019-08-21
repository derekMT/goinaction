package main

import (
  "log"
  "os"
   _ "github.com/derekmt/goinaction/matchers"
  "github.com/derekmt/search"
)
// init is called prior to main
func init() {
    log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
    search.Run("president")
}