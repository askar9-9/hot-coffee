package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
)

var (
	InfoFilePath  = "logs/info.log"
	ErrorFilePath = "logs/error.log"
	DebugFilePath = "logs/debug.log"
)

var (
	Port int
	help bool
)

var helpText = `
Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
`

var usageTxt = `
Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show help.
  --port N     Port number.
`

func init() {
	flag.IntVar(&Port, "port", 8080, "Port number.")
	flag.BoolVar(&help, "help", false, "Show help.")

	flag.Parse()

	if help || slices.Contains(os.Args, "--help") {
		fmt.Print(helpText)
		os.Exit(0)
	}

	if len(flag.Args()) > 0 {
		log.Fatalf("Unexpected arguments: %v\n%s", flag.Args(), usageTxt)
	}

	if Port < 1024 || Port > 49151 {
		log.Fatalf("Port number must be in the range [1024, 49151]\n%s", usageTxt)
	}
}
