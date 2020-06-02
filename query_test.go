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