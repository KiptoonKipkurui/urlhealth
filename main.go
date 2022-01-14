//
// this is the entry point for our application
//
package main

import (
	"flag"
	"fmt"

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
		fmt.Errorf("encountered error", err)
	}

	var urls []string
	for _, line := range lines {

		found, err := processing.GetUrls(line)

		if err != nil {
			fmt.Errorf("encountered error", err)
		}
		urls = append(urls, found...)
	}

	// check each url for healthiness

	for _, url := range urls {

		err := httpclient.Get(url)

		if err != nil {
			fmt.Errorf("encountered error %s", err)
		}
	}
}
