package main

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseDocument_parse_tag(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("<p>Result</p>"))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil { t.Error(err) }
	q := Query{text: "text", selector: "p", index: 0}
	res := ParseDocument(doc, q)
	assert.Equal(t, "Result", res)
}

func TestParseDocument_parse_tag_attr(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("<img src=\"http://cdn.image.com\">"))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil { t.Error(err) }
	q := Query{attr: "src", selector: "img", index: 0}
	res := ParseDocument(doc, q)
	assert.Equal(t, "http://cdn.image.com", res)
}

func TestParseDocument_parse_class(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("<h1 class=\"title\">Result</h1>"))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil { t.Error(err) }
	q := Query{ selector: ".title", text: "text", index: 0}
	res := ParseDocument(doc, q)
	assert.Equal(t, "Result", res)
}

func TestParseDocument_parse_class_attr(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("<a class=\"title\" href=\"google.com\">Result</a>"))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil { t.Error(err) }
	q := Query{ selector: ".title", text: "", attr: "href", index: 0}
	res := ParseDocument(doc, q)
	assert.Equal(t, "google.com", res)
}