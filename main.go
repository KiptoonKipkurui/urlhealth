//
// this is the entry point for our application
//
package main

import (
	"flag"
	"log"

	"github.com/kiptoonkipkurui/urlhealth/files"
	"github.com/kiptoonkipkurui/urlhealth/httpclient"
	"github.com/kiptoonkipkurui/urlhealth/processing"
)

//
// main receives path to the readme file
func main() {

	path := flag.String("path", "foo", "path to the readme")
	flag.Parse()
	lines, err := files.Read(*path)

	if err != nil {
		log.Fatal(err)
	}

	var urls []string
	for _, line := range lines {

		found, err := processing.GetUrls(line)

		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, found...)
	}

	// check each url for healthiness

	for _, url := range urls {

		err := httpclient.Get(url)

		if err != nil {
			log.Fatal(err)
		}
	}
}
