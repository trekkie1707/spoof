package parsers

import (
	"regexp"
	"testing"
)

var integerTests = map[string]string{
	"i": `[0-9]{1,4}`,
	"i:1000/9999": `[0-9]{4}`,
	"i:": `[0-9]{1,4}`,
	"i:-200/-100": `-[0-9]{3}`,
}

func TestParseInt(t *testing.T) {
	for k, v := range integerTests{
		re := regexp.MustCompile(v)
		if ! re.MatchString(ParseInt(k)) {
			t.Fatalf("Test failed on item: {{" + k + "}}\nGot value: " + ParseInt(k))
		}
	}
}