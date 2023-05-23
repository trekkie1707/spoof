package parsers

import (
	"regexp"
	"testing"
)

var listTests = map[string]string{
	"l": `[a-z]{0}`,
	"l:1/2/3": `[123]{1}`,
	"l:1/2/3::2": `[123],[123]`,
	"l:1/2/3:::-": `[123]{1}`,
	"l:1/2/3::2:-": "[123]-[123]",
	"l::testdata/list-test.txt::-": `[123]{1}`,
	"l::testdata/list-test.txt:2:-": "[123]{1}-[123]{1}",
	"l:4/5/6:testdata/list-test.txt:10:-": "[123456-]{19}",
}

func TestParseList(t *testing.T) {
	for k, v := range listTests{
		re := regexp.MustCompile(v)
		if ! re.MatchString(ParseList(k)) {
			t.Fatalf("Test failed on item: {{" + k + "}}\nGot value: " + ParseList(k))
		}
	}
}