//
// package files contains helper methods to read files
package files

import (
	"bufio"
	"errors"
	"log"
	"os"
)

//
// Read takes in a path and returns the contents line by line
func Read(path string) ([]string, error) {

	log.Println("checking file availability...")
	// ensure the file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {

		log.Fatal(err)
		return nil, os.ErrNotExist
	}

	f, err := os.Open(path)

	if err != nil {

		log.Fatal(err)
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines = make([]string, 0)
	for scanner.Scan() {

		text := scanner.Text()
		lines = append(lines, text)
	}

	return lines, nil
}
