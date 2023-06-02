package parsers

import (
	"regexp"
	"testing"
)

var floatTests = map[string]string{
	"f": `[0-9]{1,4}\.[0-9]{2}`,
	"f:1000/9999": `[0-9]{4}\.[0-9]{2}`,
	"f:": `[0-9]{1,4}\.[0-9]{2}`,
	"f:-200/-100": `-[0-9]{3}\.[0-9]{2}`,
	"f::5": `[0-9]{1,4}\.[0-9]{5}`,
	"f:1000/9999:5": `[0-9]{4}\.[0-9]{5}`,
}

func TestParseFloat(t *testing.T) {
	for k, v := range floatTests{
		re := regexp.MustCompile(v)
		if ! re.MatchString(ParseFloat(k)) {
			t.Fatalf("Test failed on item: {{" + k + "}}\nGot value: " + ParseFloat(k))
		}
	}
}