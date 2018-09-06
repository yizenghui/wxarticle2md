package wxarticle2md

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Article struct {
	// Basic
	Html        string `json:"content_html"`
	Content     string `json:"content"`
	Title       string `json:"title"`
	Publishtime int64  `json:"publish_time"`

	// Others
	Images      []string `json:"images"`
	ReadContent string   `json:"read_content"`
	contentNode *html.Node
}

// html to
func ToAticle(htmlStr string) (htmlArticle string, err error) {

	n, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return
	}
	a := &Article{}

	a.contentNode = n
	a.cleanAttrs(a.contentNode, "class", "id", "style", "width", "height", "onclick", "onmouseover", "border")
	a.clean(a.contentNode, func(n *html.Node) bool {
		return n.Type == html.CommentNode || n.DataAtom == atom.Script || n.DataAtom == atom.Object
	})
	a.Content = getText(n)
	a.Html, err = getHtml(n)

	getImages(a.contentNode)
	htmlArticle, err = getHtml(a.contentNode)

	if err != nil {
		return
	}
	//  getImages(n)
	return
}

// ParseReadContent parse the ReadContent to be readability
func (a *Article) ParseReadContent() {
	a.cleanAttrs(a.contentNode, "class", "id", "style", "width", "height", "onclick", "onmouseover", "border")
	a.clean(a.contentNode, func(n *html.Node) bool {
		return n.Type == html.CommentNode || n.DataAtom == atom.Script || n.DataAtom == atom.Object
	})
	a.ReadContent, _ = getHtml(a.contentNode)
	// a.ReadContent = CompressHtml(a.ReadContent)
}

func (a *Article) cleanAttrs(sel *html.Node, attrs ...string) {
	for _, attr := range attrs {
		removeAttr(sel, attr)
	}
	for c := sel.FirstChild; c != nil; c = c.NextSibling {
		a.cleanAttrs(c, attrs...)
	}
}
func (a *Article) clean(sel *html.Node, toClean selector) {
	for c := sel.FirstChild; c != nil; c = c.NextSibling {
		if toClean(c) {
			pre := c.PrevSibling
			sel.RemoveChild(c)
			c = pre
		} else {
			a.clean(c, toClean)
		}
		if c == nil {
			c = sel.FirstChild
			if c == nil {
				break
			}
		}
	}
}

func (a *Article) GetContentNode() *html.Node {
	return a.contentNode
}

// html to
func GetHtml(urlStr string) (htmlArticle string, err error) {
	g, e := goquery.NewDocument(urlStr)
	if e != nil {
	}

	contentHTML, err := g.Find("#js_content").Html()
	if err != nil {
	}
	html2 := fmt.Sprintf(`
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<title>%v</title>
		<body><div class="rich_media_content " id="js_content">
		%v
		</div>
		</body>
		</html>
		`, `NONE TITLE`, contentHTML)

	return html2, err
}
