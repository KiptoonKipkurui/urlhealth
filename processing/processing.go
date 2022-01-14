//
// pacakge processing contains functionality to determine if a line of file contains a url
//
package processing

import "regexp"

//
// get urls from a line on the file
//
func GetUrls(line string) ([]string, error) {

	re := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)`)

	found := re.FindAllString(line, -1)

	return found, nil
}
