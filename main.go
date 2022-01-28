//
// this is the entry point for our application
//
package main

import (
	"flag"
	"fmt"
	"sync"

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

	wg := &sync.WaitGroup{}
	for _, readme := range result {
		for _, link := range readme.Links {
			wg.Add(1)
			go func(link files.Link, wg *sync.WaitGroup) {

				err := httpclient.Get(link.Url, wg)
				if err != nil {
					fmt.Printf("ecountered error %s for link %s", err.Error(), link.Url)
				} else {
					link.Healthy = true
				}
			}(link, wg)
		}
	}

	wg.Wait()
}
