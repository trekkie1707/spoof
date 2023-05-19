package parsers

import (
	"regexp"
	"testing"
)

var booleanTests = map[string]string{
	"b": `true|false`,
	"b:true": `True|False`,
	"b:asdf": `True|False`,
	"b:": `True|False`,
}

func TestParseBoolean(t *testing.T) {
	for k, v := range booleanTests{
		re := regexp.MustCompile(v)
		if ! re.MatchString(ParseBool(k)) {
			t.Fatalf("Test failed on item: {{" + k + "}}\nGot value: " + ParseBool(k))
		}
	}
}