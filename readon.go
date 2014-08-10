package readon

import (
	"code.google.com/p/go-html-transform/h5"
	"code.google.com/p/go-html-transform/html/transform"
	"code.google.com/p/go.net/html"
	"io"
	"math"
	"strings"
)

type Article struct {
	Title       string
	ArticleHtml string
}

func NewArticle(reader io.Reader) (*Article, error) {
	tree, _ := h5.New(reader)
	t := transform.New(tree)
	removeScripts(t)
	removeUnlikely(t)
	removeCss(t)
	removeImages(t)
	// TODO(pascalj): replace double BR with P
	removeBr(t)
	removeTags(t, []string{"form", "h1", "object", "iframe"})
	removeEmpty(t)
	topTag := topCandidate(t.Doc())
	article := &Article{"", h5.RenderNodesToString([]*html.Node{topTag})}
	return article, nil
}

func removeScripts(t *transform.Transformer) {
	t.Apply(transform.Replace(), "script")
	t.Apply(transform.Replace(), "noscript")
}

func removeCss(t *transform.Transformer) {
	t.Apply(transform.Replace(), "link")
	t.Apply(transform.ModifyAttrib("style", ""), "[style]")
}

func removeImages(t *transform.Transformer) {
	t.Apply(transform.Replace(), "figure")
	t.Apply(transform.Replace(), "img")
}

func removeUnlikely(t *transform.Transformer) {
	unlikelyClasses := ".combx, .comment, .community, .disqus, .extra, .foot, .header, .menu, .remark, .rss, .shoutbox, .sidebar, .sponsor, .ad-break, .agegate, .pagination, .pager, .popup, .tweet, .twitter, .ad"
	unlikelyIds := "#combx, #comment, #community, #disqus, #extra, #foot, #header, #menu, #remark, #rss, #shoutbox, #sidebar, #sponsor, #ad-break, #agegate, #pagination, #pager, #popup, #tweet, #twitter"
	applyGroup(unlikelyClasses, func(sel string) { t.Apply(transform.Replace(), sel) })
	applyGroup(unlikelyIds, func(sel string) { t.Apply(transform.Replace(), sel) })
}

func removeBr(t *transform.Transformer) {
	t.Apply(transform.Replace(), "br")
}

func removeEmpty(t *transform.Transformer) {
	t.Apply(transform.Replace(), "li:empty")
	t.Apply(transform.Replace(), "p:empty")
}

func removeTags(t *transform.Transformer, tags []string) {
	for _, tag := range tags {
		t.Apply(transform.Replace(), tag)
	}
}

func topCandidate(node *html.Node) *html.Node {
	ratings := rateCandidates(node)
	var topCandidate *html.Node

	// Weight the links by linkDensity (less links is better) and get the top candidate
	for node, score := range ratings {
		if topCandidate == nil || score > ratings[topCandidate] {
			topCandidate = node
		}
	}

	return topCandidate
}

func rateCandidates(node *html.Node) map[*html.Node]int {
	ratings := make(map[*html.Node]int)
	h5.WalkNodes(node, func(node *html.Node) {

		// Only consider paragraphs
		if node.Type != html.ElementNode || node.Data != "p" {
			return
		}
		score := rateNode(node)
		if node.Parent != nil {
			if _, containsNode := ratings[node.Parent]; !containsNode {
				ratings[node.Parent] = 0
			}
			ratings[node.Parent] += int(float32(score) * (1.0 - linkDensity(node.Parent)))
		}
		if node.Parent != nil && node.Parent.Parent != nil {
			if _, containsNode := ratings[node.Parent.Parent]; !containsNode {
				ratings[node.Parent.Parent] = 0
			}
			ratings[node.Parent.Parent] += int(float32(score)*(1.0-linkDensity(node.Parent.Parent))) / 2
		}
	})
	return ratings
}

func rateNode(node *html.Node) int {
	score := 1
	innerText := innerText(node)

	// Don't even consider short paragraphs
	if len(innerText) <= 25 {
		return 0
	}

	// Add a point for ervery comma
	score += strings.Count(innerText, ",")

	// Add up to three points for each 100 chars
	if hChars := int(math.Floor(float64(len(innerText)) / 100)); hChars <= 3 {
		score += hChars
	} else {
		score += 3
	}
	return score
}

func linkDensity(node *html.Node) float32 {
	var textLength, linkLength int
	textLength = len(innerText(node))
	h5.WalkNodes(node, func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			linkLength = linkLength + len(innerText(node))
		}
	})
	return float32(linkLength) / float32(textLength)
}

func innerText(node *html.Node) string {
	var content string
	h5.WalkNodes(node, func(node *html.Node) {
		if node.Type == html.TextNode {
			content = content + node.Data
		}
	})
	return content
}

func count(node *html.Node, tag string) int {
	var total int
	h5.WalkNodes(node, func(node *html.Node) {
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
