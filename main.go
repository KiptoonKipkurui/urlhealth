//
// this is the entry point for our application
//
package main

import (
	"flag"
	"fmt"

	"github.com/kiptoonkipkurui/urlhealth/files"
	"github.com/kiptoonkipkurui/urlhealth/httpclient"
)

//
// main receives path to the readme file
func main() {

	path := flag.String("path", "foo", "path to the readme")
	flag.Parse()
	readmes := make([]files.ReadMe, 0)
	result, err := files.Discover(*path, readmes)

	fmt.Println(result)

	if err != nil {
		fmt.Printf("ecountered error %s", err.Error())
	}

	for _, readme := range result {
		for _, link := range readme.Links {

			err := httpclient.Get(link.Url)
			if err != nil {
				fmt.Printf("ecountered error %s for link %s", err.Error(), link.Url)
			}
		}
	}
}
