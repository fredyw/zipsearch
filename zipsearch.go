// Copyright 2013 Fredy Wijaya
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func validateArgs() {
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:", os.Args[0], "<directory> <regex_pattern>")
}

func readZipFile(zipFile string, regex *regexp.Regexp) error {
	r, e := zip.OpenReader(zipFile)
	if e != nil {
		return e
	}
	defer r.Close()

	for _, f := range r.File {
		matched := regex.MatchString(f.Name)
		if matched {
			fmt.Println("Found", f.Name, "in", zipFile)
		}
	}
	return nil
}

func main() {
	validateArgs()
	regex, e := regexp.Compile(os.Args[2])
	if e != nil {
		fmt.Errorf("Error: invalid regular expression: %s", os.Args[2])
		os.Exit(1)
	}
	filepath.Walk(os.Args[1],
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				absPath, e := filepath.Abs(path)
				if e != nil {
					return e
				}
				readZipFile(absPath, regex)
			}
			return nil
		})
}
