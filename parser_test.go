package main

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func CreateDoc(str string) *goquery.Document {
	reader := bufio.NewReader(strings.NewReader(strings.TrimSpace(str)))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil { log.Fatal(err) }
	return doc
}

func TestParseDocument_parse_tag(t *testing.T) {
	q := Query{text: "text", selector: "p", index: 0}
	doc := CreateDoc("<p>Result</p>")
	res := ParseDocument(doc, q)
	assert.Equal(t, "Result", res)
}

func TestParseDocument_parse_tag_attr(t *testing.T) {
	q := Query{attr: "src", selector: "img", index: 0}
	doc := CreateDoc("<img src=\"http://cdn.image.com\">")
	res := ParseDocument(doc, q)
	assert.Equal(t, "http://cdn.image.com", res)
}

func TestParseDocument_parse_class(t *testing.T) {
	q := Query{ selector: ".title", text: "text", index: 0}
	doc := CreateDoc("<h1 class=\"title\">Result</h1>")
	res := ParseDocument(doc, q)
	assert.Equal(t, "Result", res)
}

func TestParseDocument_parse_class_attr(t *testing.T) {
	q := Query{ selector: ".title", text: "", attr: "href", index: 0}
	doc := CreateDoc("<a class=\"title\" href=\"google.com\"></a>")
	res := ParseDocument(doc, q)
	assert.Equal(t, "google.com", res)
}

func TestParseDocument_parse_class_attr_index(t *testing.T) {
	q := Query{ selector: ".title", text: "", attr: "href", index: 0}
	doc := CreateDoc(`
		<a class="title"  href="google.com"></a>
	`)
	res := ParseDocument(doc, q)
	assert.Equal(t, "google.com", res)
}