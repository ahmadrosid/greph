package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseQuery(t *testing.T) {
	query, _ := ParseQuery("p[0].text")
	assert.Equal(t, "p", query.selector)
	assert.Equal(t, 0, query.index)
	assert.Equal(t, "text", query.text)

	query, _ = ParseQuery("h1[0].text")
	assert.Equal(t, "h1", query.selector)
	assert.Equal(t, 0, query.index)
	assert.Equal(t, "text", query.text)

	query, _ = ParseQuery(".container[1].text")
	assert.Equal(t, ".container", query.selector)
	assert.Equal(t, 1, query.index)
	assert.Equal(t, "text", query.text)
}

func TestParseQueryIndex(t *testing.T) {
	query, _ := ParseQuery("p[0:5].text")
	assert.Equal(t, "p", query.selector)
	assert.Equal(t, -1, query.index)
	assert.Equal(t, 0, query.indexStart)
	assert.Equal(t, 5, query.indexEnd)
	assert.Equal(t, "text", query.text)

	query, _ = ParseQuery("p[1:5].text")
	assert.Equal(t, "p", query.selector)
	assert.Equal(t, -1, query.index)
	assert.Equal(t, 1, query.indexStart)
	assert.Equal(t, 5, query.indexEnd)
	assert.Equal(t, "text", query.text)

	query, _ = ParseQuery("p[1:].text")
	assert.Equal(t, "p", query.selector)
	assert.Equal(t, -1, query.index)
	assert.Equal(t, 1, query.indexStart)
	assert.Equal(t, 0, query.indexEnd)
	assert.Equal(t, "text", query.text)

	query, _ = ParseQuery("p[:3].text")
	assert.Equal(t, "p", query.selector)
	assert.Equal(t, -1, query.index)
	assert.Equal(t, 0, query.indexStart)
	assert.Equal(t, 3, query.indexEnd)
	assert.Equal(t, "text", query.text)

	query, _ = ParseQuery("p[:].text")
	assert.Equal(t, "p", query.selector)
	assert.Equal(t, -1, query.index)
	assert.Equal(t, 0, query.indexStart)
	assert.Equal(t, 0, query.indexEnd)
	assert.Equal(t, "text", query.text)
}