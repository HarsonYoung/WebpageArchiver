package common

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"regexp"
	"strings"
)

func ExtractHTMLText(htmlContent string) (string, error) {
	r := strings.NewReader(htmlContent)
	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}

	var text bytes.Buffer
	skip := false

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script") {
			skip = true
		} else if n.Type == html.TextNode && !skip {
			text.WriteString(strings.TrimSpace(n.Data) + " ")
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
			if n.Type == html.ElementNode && (n.Data == "style" || n.Data == "script") {
				skip = false
			}
		}
	}

	f(doc)

	pureHTML := text.String()
	pureHTML = strings.ReplaceAll(pureHTML, " ", "")
	pureHTML = strings.ReplaceAll(pureHTML, "\n", "")

	return pureHTML, nil
}

func GetHTMLTitle(htmlContent string) (title string, err error) {
	re := regexp.MustCompile(`(?i)<title>(.*?)</title>`)
	matches := re.FindStringSubmatch(htmlContent)
	if len(matches) < 2 {
		return "", fmt.Errorf("no title found")
	}
	return matches[1], nil
}

func GetHTMLFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
