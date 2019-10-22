package main

import (
	"fmt"
	"io/ioutil"
	"os"

	termutil "github.com/andrew-d/go-termutil"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/ini.v1"
)

var (
	inFile = kingpin.Arg("file", "INI file.").String()
	// pretty = kingpin.Flag("pretty", "Pretty print result.").Short('p').Bool()
	quiet = kingpin.Flag("quiet", "Don't output on success.").Short('q').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	data, err := readPipeOrFile(*inFile)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	filename := "-"
	if *inFile != "" {
		filename = *inFile
	}

	_, err = ini.Load(data)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	if !*quiet {
		fmt.Println("OK:", filename)
	}
}

// readPipeOrFile reads from stdin if pipe exists, else from provided file
func readPipeOrFile(fileName string) ([]byte, error) {
	if !termutil.Isatty(os.Stdin.Fd()) {
		return ioutil.ReadAll(os.Stdin)
	}
	if fileName == "" {
		return nil, fmt.Errorf("no piped data and no file provided")
	}
	return ioutil.ReadFile(fileName)
}
