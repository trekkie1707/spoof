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
)

var fileFlag = flag.Bool("f", false, "Input will be read from the file provided")

var exprBroad = regexp.MustCompile(`{{.*}}`)
var exprEnd = regexp.MustCompile(`{{[^}]*\x00`)
var exprCutoff = regexp.MustCompile(`{{[^}]*`)
var expr = regexp.MustCompile(`{{[^{}]+}}`)

func parse(value string) string {
	ret := strings.Replace(value, value, value[2:len(value)-2], 1)
	return ret
}

func parseInput(input string) {
	for exprBroad.MatchString(input) {
		for _, match := range expr.FindAllString(input, -1) {
			input = strings.Replace(input, match, parse(match), 1)
		}
	}
	fmt.Print(input)
}

func read(r *bufio.Reader) ([]byte, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return ln, err
}

func parseFiles(files ...string) {
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
			// fmt.Println(string(line))
			read = num
			fileIndex += num
			indexes := exprCutoff.FindAllIndex(line, -1)
			if (indexes != nil && indexes[len(indexes)-1][1] == len(line)) || exprEnd.Match(line) {
				var index int
				if indexes != nil {
					index = indexes[len(indexes)-1][0]
					// fmt.Println(index)
					line = line[0:index]
					// fmt.Println("")
					// fmt.Println(num)
					// fmt.Println(file.Seek(0, os.SEEK_CUR))
					file.Seek(int64(fileIndex), os.SEEK_SET)
					// fmt.Println(file.Seek(0, os.SEEK_CUR))
					file.Seek(int64(index-num), os.SEEK_CUR)
					fileIndex += index - num
					// fmt.Println(file.Seek(0, os.SEEK_CUR))
					reader.Reset(file)
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("a real error happened here: %v\n", err)
			}
			parseInput(string(line))
		}
	}
}

func main() {
	flag.Parse()
	if *fileFlag {
		parseFiles(flag.Args()...)
	} else {
		parseInput(strings.Join(flag.Args(), " "))
	}
}
