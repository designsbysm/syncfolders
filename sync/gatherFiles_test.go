package sync

import (
	"regexp"
	"testing"
)

func TestRegEx(t *testing.T) {
	exemplars := []string{
		"pattern",
		"a/pattern",
		"pattern/a",
		"a/pattern/a",
	}

	pattern := "pattern"
	re := regexp.MustCompile(pattern)
	for _, exemplar := range exemplars {
		t.Log(pattern, exemplar, re.MatchString(exemplar))
	}
	t.Log("")

	pattern = ".*pattern.*"
	re = regexp.MustCompile(pattern)
	for _, exemplar := range exemplars {
		t.Log(pattern, exemplar, re.MatchString(exemplar))
	}
	t.Log("")

	pattern = "^pattern"
	re = regexp.MustCompile(pattern)
	for _, exemplar := range exemplars {
		t.Log(pattern, exemplar, re.MatchString(exemplar))
	}
	t.Log("")

	pattern = "pattern$"
	re = regexp.MustCompile(pattern)
	for _, exemplar := range exemplars {
		t.Log(pattern, exemplar, re.MatchString(exemplar))
	}
	t.Log("")

	pattern = "^pattern$"
	re = regexp.MustCompile(pattern)
	for _, exemplar := range exemplars {
		t.Log(pattern, exemplar, re.MatchString(exemplar))
	}

	t.Fatal("")
}
