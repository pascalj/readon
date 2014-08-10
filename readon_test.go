package readon

import (
	"os"
	"regexp"
	"testing"
)

func TestRemovesScriptTags(t *testing.T) {
	file, _ := os.Open("fixtures/heise.html")
	output, _ := Readon(file)
	re := regexp.MustCompile(`<script`)

	if re.MatchString(output) {
		t.Error("Script tags were not removed")
	}
}

func TestRemovesComments(t *testing.T) {
	file, _ := os.Open("fixtures/heise.html")
	output, _ := Readon(file)
	re := regexp.MustCompile(`class="rss"`)
	if re.MatchString(output) {
		t.Error("rss class was not removed")
	}
}

func TestRemovesBr(t *testing.T) {
	file, _ := os.Open("fixtures/heise.html")
	output, _ := Readon(file)
	re := regexp.MustCompile(`<br/?>`)
	if re.MatchString(output) {
		t.Error("br tags were not removed")
	}
}
