package main

import (
	"fmt"
	"regexp"
	"strings"
)

var versionHeader = map[string]string{
	"HTML5":                  `<!doctype html>`,
	"HTML 4.01 Strict":       `<!doctype html public "-//w3c//dtd html 4.01//en"`,
	"HTML 4.01 Transitional": `<!doctype html public "-//w3c//dtd html 4.01 transitional//en"`,
	"HTML 4.01 Frameset":     `<!doctype html public "-//w3c//dtd html 4.01 frameset//en"`,
	"HTML 3.2":               `<!doctype html public "-//w3c//dtd html 3.2 final//en"`,
	"HTML 2.0":               `<!doctype html public "-//ietf//dtd html 2.0//en"`,
	"XHTML 1.0 Strict":       `<!doctype html public "-//w3c//dtd xhtml 1.0 strict//en"`,
	"XHTML 1.0 Transitional": `<!doctype html public "-//w3c//dtd xhtml 1.0 transitional//en"`,
	"XHTML 1.0 Frameset":     `<!doctype html public "-//w3c//dtd xhtml 1.0 frameset//en"`,
	"XHTML 1.1":              `<!doctype html public "-//w3c//dtd xhtml 1.1//en"`,
	"XHTML Basic 1.1":        `<!doctype html public "-//w3c//dtd xhtml basic 1.1//en"`,
	"XHTML Basic 1.0":        `<!doctype html public "-//w3c//dtd xhtml basic 1.0//en"`,
}

type RegexPageAnalyser struct {
}

func (r *RegexPageAnalyser) Analyse(page string) (PageInfo, error) {
	page = strings.ReplaceAll(strings.ToLower(page), "\n", " ")
	links := links(page)
	internal := countInternal(links)
	external := len(links) - internal
	inaccessible := countInaccessible(links)
	pageInfo := PageInfo{
		HtmlVersion:       version(page),
		PageTitle:         title(page),
		Headers:           countHeaders(page),
		InternalLinks:     internal,
		ExternalLinks:     external,
		InaccessibleLinks: inaccessible,
		HasLoginForm:      false,
		Content:           page,
	}
	return pageInfo, nil
}

type link string

// - http://example.org/about
// - https://example.org/about
// - /about
// - //example.org/about
// - //www.example.org/about
// - #about
func (l *link) isInternal() bool {
	lnk := string(*l)
	return strings.HasPrefix(lnk, "#") ||
		strings.HasPrefix(lnk, "/")
}

func links(page string) []link {
	href := regexp.MustCompile(`<a.*?href="(.*?)".*?>.*?</a>`)
	hrefMatch := href.FindAllStringSubmatch(page, -1)
	links := make([]link, 0)
	for _, match := range hrefMatch {
		if len(match) > 0 {
			links = append(links, link(match[1]))
		}
	}
	return links
}

func countInternal(links []link) int {
	count := 0
	for _, v := range links {
		if v.isInternal() {
			count++
		}
	}
	return count
}

func countInaccessible(links []link) int {
	return 0
}

var headers = []string{"h1", "h2", "h3", "h4", "h5", "h6"}

func countHeaders(page string) map[string]int {
	result := make(map[string]int, len(headers))
	for _, header := range headers {
		result[header] = strings.Count(page, fmt.Sprintf("</%s>", header))
	}
	return result
}

func title(page string) string {
	title := regexp.MustCompile("<title>(.*)</title>")
	titleMatch := title.FindAllStringSubmatch(page, -1)

	if len(titleMatch) > 0 && len(titleMatch[0]) > 0 {
		return titleMatch[0][1]
	}
	return ""
}

func version(page string) string {
	for k, v := range versionHeader {
		if strings.HasPrefix(page, v) {
			return k
		}
	}
	return ""
}
