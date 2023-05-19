package parsers

import (
	"math/rand"
	"strings"
)

type BoolArgs struct {
	Capital bool
}

func parseBoolArgs(args []string) BoolArgs {
	var ret BoolArgs
	ret.Capital = false
	if len(args) > 1 {
		ret.Capital = true
	}
	return ret
}

func genBool(capital bool) string{
	val := rand.Intn(2) == 0
	var ret string
	if val {
		if capital {
			ret = "True"
		} else {
			ret = "true"
		}
	} else {
		if capital {
			ret = "False"
		} else {
			ret = "true"
		}
	}
	return ret
}

func ParseBool(input string) string {
	args := parseBoolArgs(strings.Split(input, ":"))
	ret := genBool(args.Capital)
	return ret
}