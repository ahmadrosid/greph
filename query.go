package main

import (
	"strconv"
	"strings"
)

type Query struct {
	selector string
	attr string
	text string
	index int
}

func ParseQuery(s string) (Query, error) {
	q := Query{}
	q.index = -1
	key := ""

	for i := 0; i < len(s); i++ {
		if i == 0 {
			next := strings.IndexByte(s,'[')
			q.selector = s[:next]
			i = next
		}
		if s[i] == '[' {
			next := strings.IndexByte(s,']')
			indexStr := s[i+1:next]
			index, err := strconv.Atoi(indexStr)
			q.index = index
			if err != nil { q.index = -1 }
			i = next
		}
		if q.index != -1 {
			key = s[i+1:]
			i = len(s)
		}
	}

	if strings.Contains(key, ":") {
		q.attr = key[1:]
	} else {
		q.text = key[1:]
	}

	return q, nil
}
