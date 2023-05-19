package parsers

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type IntArgs struct {
	Min int
	Max int
}

func parseIntArgs(args []string) IntArgs {
	var ret IntArgs
	ret.Min = 0
	ret.Max = 1000
	if len(args) > 1 {
		minMaxStr := strings.Split(args[1],"/")
		if len(minMaxStr) > 1{
			min, minErr := strconv.Atoi(minMaxStr[0])
			max, maxErr := strconv.Atoi(minMaxStr[1])
			if minErr != nil || maxErr != nil {
				log.Fatalf("Integer parsing failed on token: {{%s}}", strings.Join(args,":"))
			}
			ret.Min = min
			ret.Max = max
		}
	}
	return ret
}

func genInt(min int, max int) string{
	return fmt.Sprint(rand.Intn(max - min) + min)
}

func ParseInt(input string) string {
	args := parseIntArgs(strings.Split(input, ":"))
	ret := genInt(args.Min, args.Max)
	return ret
}