package parsers

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/trekkie1707/strl/constants"
)

type StringArgs struct {
	Length int
	Pool []string
}

func parseStringArgs(args []string) StringArgs {
	var ret StringArgs
	ret.Length = 5
	// fmt.Println(len(args))
	if len(args) > 1 {
		intVal, err := strconv.Atoi(args[1])
		if err == nil && intVal > 0{
			ret.Length = intVal
		}
	}
	if len(args) > 2 {
		typeStr := args[2]
		// fmt.Println(strings.Index(typeStr, "a"))
		ret.Pool = make([]string,0)
		if strings.Contains(typeStr, "a") {
			ret.Pool = append(ret.Pool, constants.GetLowerPool()...)
		}
		if strings.Contains(typeStr, "A") {
			ret.Pool = append(ret.Pool, constants.GetUpperPool()...)
		}
		if strings.Contains(typeStr, "0") {
			ret.Pool = append(ret.Pool, constants.GetNumberPool()...)
		}
		if strings.Contains(typeStr, "?") {
			ret.Pool = append(ret.Pool, constants.GetSymbolPool()...)
		}
		if len(ret.Pool) == 0 {
			ret.Pool = append(ret.Pool, constants.GetAllPool()...)
		}
	} else {
		ret.Pool = append(ret.Pool, constants.GetAllPool()...)
	}
	// fmt.Println(ret)
	return ret
}

func genString(length int, pool []string) string{
	ret := ""
	// fmt.Println(len(pool))
	max := len(pool)
	for length > 0 {
		ret += pool[rand.Intn(max)]
		length -= 1
	}
	return ret
}

func ParseString(input string) string {
	args := parseStringArgs(strings.Split(input, ":"))
	ret := genString(args.Length, args.Pool)
	return ret
}