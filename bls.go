// (c) copyright 2013. Johannes Boyne
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Reset      = "\x1b[0m"
	Bright     = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underscore = "\x1b[4m"

	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"

	BgBlack = "\x1b[40m"
	BgWhite = "\x1b[47m"
)

func main() {
	// path checking
	var path string
	if len(os.Args) <= 1 {
		path = "."
	} else {
		path = os.Args[1]
	}

	// file counter / information aggregation
	dir, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	// pseudo table
	type Table struct {
		LongestName int
		LongestSize int
		LongestMode int
	}

	t := Table{0, 0, 0}

	for _, fi := range fileInfos {
		if len(fi.Name()) > t.LongestName {
			t.LongestName = len(fi.Name())
		}
		if len(strconv.FormatInt(fi.Size(), 10)) > t.LongestSize {
			t.LongestSize = len(strconv.FormatInt(fi.Size(), 10))
		}
		if len(fi.Mode().String()) > t.LongestMode {
			t.LongestMode = len(fi.Mode().String())
		}
	}
	// list filenames
	for _, fi := range fileInfos {
		fmt.Println(
			fi.Name(),
			strings.Repeat(" ", t.LongestName-len(fi.Name())),
			FgGreen+strconv.FormatInt(fi.Size(), 10)+Reset,
			strings.Repeat(" ", t.LongestSize-len(strconv.FormatInt(fi.Size(), 10))),
			fi.Mode())
	}
	// print it
	fmt.Println(strings.Repeat("-", t.LongestSize+t.LongestName+t.LongestMode) + "----")
	fmt.Println(len(fileInfos))
}
