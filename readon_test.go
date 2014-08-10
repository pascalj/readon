package readon

import (
	"os"
	"regexp"
	"testing"
)

var testArticle *Article

func setup() {
	if testArticle == nil {
		file, _ := os.Open("fixtures/heise.html")
		testArticle, _ = NewArticle(file)
	}
}

func TestRemovesScriptTags(t *testing.T) {
	setup()
	re := regexp.MustCompile(`<script`)
	if re.MatchString(testArticle.ArticleHtml) {
		t.Error("Script tags were not removed")
	}
}

func TestRemovesComments(t *testing.T) {
	setup()
	re := regexp.MustCompile(`class="rss"`)
	if re.MatchString(testArticle.ArticleHtml) {
		t.Error("rss class was not removed")
	}
}

func TestRemovesBr(t *testing.T) {
	setup()
	re := regexp.MustCompile(`<br/?>`)
	if re.MatchString(testArticle.ArticleHtml) {
		t.Error("br tags were not removed")
	}
}

func TestFindsH1Title(t *testing.T) {
	setup()
	expected := "Sensorfußboden erkennt Einbrecher"
	if testArticle.Title != expected {
		t.Error("The title was not found")
	}
}
