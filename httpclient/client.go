//
// package httpclient contains functionality to make http calls
//
package httpclient

import (
	"fmt"
	"log"
	"net/http"
)

//
// Get makes a http call and returns if any error is encountered
//
func Get(url string) error {

	// make http call

	log.Printf("making a call to url %s", url)
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("ecountered error %s", err.Error())
		return err
	}

	log.Printf("obtained response status %s", response.Status)

	return nil
}
