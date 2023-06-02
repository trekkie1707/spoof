package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	myParsers "github.com/trekkie1707/strl/parsers"
)

var fileFlag = flag.Bool("f", false, "Input will be read from the file provided")

var exprBroad = regexp.MustCompile(`{{.*}}`)
var exprEnd = regexp.MustCompile(`{{[^}]*\x00`)
var exprCutoff = regexp.MustCompile(`{{[^}]*`)
var expr = regexp.MustCompile(`{{[^{}]+}}`)

var parsers = map[string]func(string)string{
	"s": myParsers.ParseString,
	"b": myParsers.ParseBool,
	"i": myParsers.ParseInt,
	"l": myParsers.ParseList,
	"f": myParsers.ParseFloat,
}

func parse(value string) string {
	ret := strings.Replace(value, value, value[2:len(value)-2], 1)
	function, found := parsers[string(ret[0])]
	if found {
		ret = function(ret)
	}
	return ret
}

func ParseInput(input string) {
	for exprBroad.MatchString(input) {
		for _, match := range expr.FindAllString(input, -1) {
			fmt.Printf("Before: %s\n", match)
			input = strings.Replace(input, match, parse(match), 1)
		}
	}
	fmt.Print(input)
}

func ParseFiles(files ...string) {
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("could not open the file: %v", err)
		}
		reader := bufio.NewReader(file)
		read := 1
		fileIndex := 0
		for read > 0 {
			line := make([]byte, 1024)
			num, err := reader.Read(line)
			read = num
			fileIndex += num
			indexes := exprCutoff.FindAllIndex(line, -1)
			if (indexes != nil && indexes[len(indexes)-1][1] == len(line)) || exprEnd.Match(line) {
				var index int
				if indexes != nil {
					index = indexes[len(indexes)-1][0]
					line = line[0:index]
					file.Seek(int64(fileIndex), io.SeekCurrent)
					file.Seek(int64(index-num), io.SeekCurrent)
					fileIndex += index - num
					reader.Reset(file)
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("a real error happened here: %v\n", err)
			}
			ParseInput(string(line))
		}
	}
}

func main() {
	flag.Parse()
	if *fileFlag {
		ParseFiles(flag.Args()...)
	} else {
		ParseInput(strings.Join(flag.Args(), " "))
	}
}
