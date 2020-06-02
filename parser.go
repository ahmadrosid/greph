package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ParseDocument(doc *goquery.Document, query Query) string {
	var result string
	doc.Find(query.selector).Each(func(i int, selection *goquery.Selection) {
		if query.index == i {
			if query.text != "" {
				result += selection.Text()
			}

			if query.attr != "" {
				attr, exists := selection.Attr(query.attr)
				if exists { result += attr }
			}
		}
	})

	return strings.TrimSpace(result)
}