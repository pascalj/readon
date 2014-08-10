package readon

import (
	"os"
	"regexp"
	"testing"
)

func TestRemovesScriptTags(t *testing.T) {
	file, _ := os.Open("fixtures/heise.html")
	article, _ := NewArticle(file)
	re := regexp.MustCompile(`<script`)

	if re.MatchString(article.ArticleHtml) {
		t.Error("Script tags were not removed")
	}
}

func TestRemovesComments(t *testing.T) {
	file, _ := os.Open("fixtures/heise.html")
	article, _ := NewArticle(file)
	re := regexp.MustCompile(`class="rss"`)
	if re.MatchString(article.ArticleHtml) {
		t.Error("rss class was not removed")
	}
}

func TestRemovesBr(t *testing.T) {
	file, _ := os.Open("fixtures/heise.html")
	article, _ := NewArticle(file)
	re := regexp.MustCompile(`<br/?>`)
	if re.MatchString(article.ArticleHtml) {
		t.Error("br tags were not removed")
	}
}
