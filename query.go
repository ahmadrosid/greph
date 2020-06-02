package main

import (
	"strconv"
	"strings"
)

type Query struct {
	selector   string
	attr       string
	text       string
	index      int
	indexStart int
	indexEnd   int
}

func (q *Query) ParseIndex(i *int, s string) {
	next := strings.IndexByte(s,']')
	indexStr := s[*i+1:next]
	index, err := strconv.Atoi(indexStr)
	q.index = index
	if err != nil { q.index = -1 }
	if strings.Contains(indexStr, ":") {
		startIndex := strings.IndexByte(indexStr,':')
		startIndexStr := indexStr[:startIndex]
		parseIndex, err := strconv.Atoi(startIndexStr)
		q.indexStart = parseIndex
		if err != nil {
			q.indexStart = 0
		}
		nextIndexStr := indexStr[startIndex+1:]
		parseIndex, err = strconv.Atoi(nextIndexStr)
		q.indexEnd = parseIndex
		if err != nil {
			q.indexEnd = 0
		}
	}
	*i = next
}

func (q *Query) ParseSelector(i *int, s string) {
	next := strings.IndexByte(s,'[')
	q.selector = s[:next]
	*i = next
}

func (q *Query) ParseExtractor(i *int, s string) {
	key := s[*i+1:]
	if strings.Contains(key, ":") {
		q.attr = key[1:]
	} else {
		q.text = key[1:]
	}
	*i = len(s)
}

func (q *Query) MatchIndex(i int) bool {
	One := q.index != -1                           // [0]
	All := q.indexStart == 0 && q.indexEnd == 0    // [:]
	Between := q.indexStart > 0 && q.indexEnd > 0  // [2:5]
	AllFrom := q.indexStart > 0 && q.indexEnd == 0 // [2:]
	Until := q.indexStart == 0 && q.indexEnd > 0   // [:5]

	if One {
		return i == q.index
	}

	if All {
		return true
	}

	if Between {
		return i >= q.indexStart && i <= q.indexEnd
	}

	if AllFrom {
		return i >= q.indexStart
	}

	if Until {
		return i <= q.indexEnd
	}

	return false
}

func ParseQuery(s string) (Query, error) {
	q := Query{}
	q.index = -1
	q.indexStart = -1
	q.indexEnd = -1

	for i := 0; i < len(s); i++ {
		if i == 0 {
			q.ParseSelector(&i, s)
		}
		if s[i] == '[' {
			q.ParseIndex(&i, s)
		}
		if q.index != -1 || q.indexStart != -1 || q.indexEnd != -1 {
			q.ParseExtractor(&i, s)
		}
	}

	return q, nil
}
