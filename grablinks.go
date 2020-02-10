package grablinks

import (
	"golang.org/x/net/html"
	"io"
	"strconv"
	"strings"
)

func All(body io.Reader) []string {
	var links []string
	var isAdded = make(map[string]bool)

	page := html.NewTokenizer(body)

	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					url := attr.Val
					if strings.Contains(url, "#") {
						url = trimHash(url)
					}
					if !isAdded[url] {
						links = append(links, url)
						isAdded[url] = true
					}
				}
			}
		}
	}
}

func trimHash(l string) string {
	var index int
	for i, str := range l {
		if strconv.QuoteRune(str) == "'#'" {
			index = i
			break
		}
	}
	return l[:index]
}
