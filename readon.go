package readon

import (
	"code.google.com/p/go-html-transform/h5"
	"code.google.com/p/go-html-transform/html/transform"
	"code.google.com/p/go.net/html"
	"fmt"
	"io"
	"strings"
)

func Readon(reader io.Reader) (string, error) {
	tree, _ := h5.New(reader)
	t := transform.New(tree)
	removeScripts(t)
	removeUnlikely(t)
	removeCss(t)
	removeBr(t)
	removeTags(t, []string{"form", "h1", "object", "iframe"})
	// TODO(pascalj): replace double BR with P
	return t.String(), nil
}

func removeScripts(t *transform.Transformer) {
	t.Apply(transform.Replace(), "script")
	t.Apply(transform.Replace(), "noscript")
}

func removeCss(t *transform.Transformer) {
	t.Apply(transform.Replace(), "link")
	t.Apply(transform.ModifyAttrib("style", ""), "[style]")
}

func removeUnlikely(t *transform.Transformer) {
	unlikelyClasses := ".combx, .comment, .community, .disqus, .extra, .foot, .header, .menu, .remark, .rss, .shoutbox, .sidebar, .sponsor, .ad-break, .agegate, .pagination, .pager, .popup, .tweet, .twitter"
	unlikelyIds := "#combx, #comment, #community, #disqus, #extra, #foot, #header, #menu, #remark, #rss, #shoutbox, #sidebar, #sponsor, #ad-break, #agegate, #pagination, #pager, #popup, #tweet, #twitter"
	applyGroup(unlikelyClasses, func(sel string) { t.Apply(transform.Replace(), sel) })
	applyGroup(unlikelyIds, func(sel string) { t.Apply(transform.Replace(), sel) })
}

func removeBr(t *transform.Transformer) {
	t.Apply(transform.Replace(), "br")
}

func removeTags(t *transform.Transformer, tags []string) {
	for _, tag := range tags {
		t.Apply(transform.Replace(), tag)
	}
}

func linkDensity(tree *h5.Tree) float32 {
	var textLength, linkLength int
	textLength = len(innerText(tree))
	h5.WalkNodes(tree.Top(), func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			subTree := h5.NewTree(node)
			linkLength = linkLength + len(innerText(&subTree))
		}
	})
	return float32(linkLength) / float32(textLength)
}

func innerText(tree *h5.Tree) string {
	var content string
	tree.Walk(func(node *html.Node) {
		if node.Type == html.TextNode {
			content = content + node.Data
		}
	})
	return content
}

func count(tree *h5.Tree, tag string) int {
	var total int
	tree.Walk(func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag {
			total = total + 1
		}
	})
	return total
}

func applyGroup(group string, applyFunc func(string)) {
	for _, sel := range strings.Split(group, ",") {
		applyFunc(strings.Trim(sel, " \t"))
	}
}
