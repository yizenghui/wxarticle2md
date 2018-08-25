package wxarticle2md

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// html to
func ToAticle(htmlStr string) (htmlArticle string, err error) {

	n, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return
	}
	getImages(n)
	htmlArticle, err = getHtml(n)

	if err != nil {
		return
	}
	//  getImages(n)
	return
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
