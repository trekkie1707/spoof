package parsers

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type FloatArgs struct {
	Min int
	Max int
	SigDigits int
}

func parseFloatArgs(args []string) FloatArgs {
	var ret FloatArgs
	ret.Min = 0
	ret.Max = 1000
	ret.SigDigits = 2
	if len(args) > 1 {
		minMaxStr := strings.Split(args[1],"/")
		if len(minMaxStr) > 1{
			min, minErr := strconv.Atoi(minMaxStr[0])
			max, maxErr := strconv.Atoi(minMaxStr[1])
			if minErr != nil || maxErr != nil {
				log.Fatalf("Float parsing failed on token: {{%s}}", strings.Join(args,":"))
			}
			ret.Min = min
			ret.Max = max
		}
	}
	if len(args) > 2 {
		sigDig, sigErr := strconv.Atoi(args[2])
		if sigErr != nil {
			log.Fatalf("Float parsing failed on token: {{%s}}", strings.Join(args,":"))
		}
		ret.SigDigits = sigDig
	}
	return ret
}

func padRight(str string, padTo int) string {
	ret := str
	padTo -= len(str) 
	for padTo > 0 {
		ret = strings.Join([]string{ret, "0"}, "")
		padTo -= 1
	}
	return ret
}

func genFloat(min int, max int, sigDigits int) string {
	decMin := int(math.Pow10(sigDigits-1))
	decMax := int(math.Pow10(sigDigits))
	decimal := padRight(genInt(decMin, decMax), sigDigits)
	number := genInt(min, max)
	return fmt.Sprintf("%s.%s", number, decimal)
}

func ParseFloat(input string) string {
	args := parseFloatArgs(strings.Split(input, ":"))
	ret := genFloat(args.Min, args.Max, args.SigDigits)
	return ret
}