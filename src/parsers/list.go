package parsers

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type ListArgs struct {
	Count int
	Pool []string
	Del string
}

func parseListArgs(args []string) ListArgs {
	var ret ListArgs
	ret.Count = 1
	ret.Del = ","
	ret.Pool = make([]string,0)
	if len(args) > 1 && len(args[1]) > 0 {
		ret.Pool = append(ret.Pool, strings.Split(args[1], "/")...)
	}
	if len(args) > 2 {
		listFile, err := os.Open(args[2])
		if err == nil {
			scanner := bufio.NewScanner(listFile)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan(){
				if len(scanner.Text()) > 0 {
					ret.Pool = append(ret.Pool, scanner.Text())
				}
			}
		} else if len(args[2]) > 0 {
			fmt.Println(err)
			fmt.Println(os.Executable())
			log.Fatalf("Error reading file {%s}", args[2])
		}
	}
	if len(args) > 3 {
		cnt, err := strconv.Atoi(args[3])
		if err == nil {
			ret.Count = cnt
		}
	}
	if len(args) > 4 {
		if len(args[4]) > 0 {
			ret.Del = args[4]	
		}
	}
	fmt.Println(ret.Pool)
	return ret
}

func genList(count int, pool []string, del string) string{
	ret := make([]string, 0)
	max := len(pool)
	if max <= 0 {
		return ""
	}
	for count > 0 {
		ret = append(ret, pool[rand.Intn(max)]) 
		count -= 1
	}
	return strings.Join(ret, del)
}

func ParseList(input string) string {
	args := parseListArgs(strings.Split(input, ":"))
	ret := genList(args.Count, args.Pool, args.Del)
	return ret
}