//change engle export csv data file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		filename = "*.csv"
	}
	infiles, _ := filepath.Glob(filename)
	os.Mkdir("out", os.ModePerm)
	for _, infile := range infiles {
		basename := filepath.Base(infile)
		outfile := filepath.Join("out", basename)
		csvUnicode(infile, outfile)
	}
}

func csvUnicode(infile, outfile string) {
	out, err := os.Create(outfile)
	if err != nil {
		os.Exit(1)
	}
	defer out.Close()
	out.WriteString("\xEF\xBB\xBF")
	//out.WriteString("\xFF\xFE\x00\x00")
	in, err := os.Open(infile)
	if err != nil {
		os.Exit(1)
	}
	c, err := ioutil.ReadAll(in)
	s := strings.Replace(string(c), ";", ",", -1)
	out.WriteString(string(s))
	fmt.Printf("Success in convert %s to %s\n", infile, outfile)
}
