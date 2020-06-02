package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		HandleStdin()
	} else {
		HandleCrawl()
	}
}

func HandleCrawl() {
	if len(os.Args) < 3 {
		fmt.Println("Please sepcify url and selector")
		fmt.Println("Usage:\n\tgreph https://example.com \"p[0].text\"")
		return
	}

	targetUrl, err := url.ParseRequestURI(os.Args[1])
	if err != nil || targetUrl.Host == "" {
		fmt.Println("Invalid url!", targetUrl)
		return
	}

	response, err := http.Get(targetUrl.String())
	StopErr(err)

	doc, err := goquery.NewDocumentFromReader(response.Body)
	StopErr(err)

	query, err := ParseQuery(os.Args[2])
	StopErr(err)

	res := ParseDocument(doc, query)
	fmt.Println(res)
}

func StopErr(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}

func HandleStdin() {
	if len(os.Args) == 1 {
		fmt.Println("Please sepcify stdin and selector")
		fmt.Println("Usage:\n\techo \"<p>Paragraph</p>\" | greph \"p[0].text\"")
		return
	}

	query, err := ParseQuery(os.Args[1])
	StopErr(err)

	doc, err := goquery.NewDocumentFromReader(bufio.NewReader(os.Stdin))
	StopErr(err)

	res := ParseDocument(doc, query)
	fmt.Println(res)
}
