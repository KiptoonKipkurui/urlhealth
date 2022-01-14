//
// package httpclient contains functionality to make http calls
//
package httpclient

import "net/http"

//
// Get makes a http call and returns if any error is encountered
//
func Get(url string) error {

	// make http call
	_, err := http.Get(url)

	if err != nil {
		return err
	}

	return nil
}
