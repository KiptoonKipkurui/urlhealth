//
// package files contains helper methods to read files
package files

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

//
// Read takes in a path and returns the contents line by line
func Read(path string) ([]string, error) {

	log.Println("checking file availability...")
	// ensure the file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {

		fmt.Printf("ecountered error %s", err.Error())
		return nil, os.ErrNotExist
	}

	log.Printf("opening file %s", path)
	f, err := os.Open(path)

	if err != nil {

		fmt.Printf("ecountered error %s", err.Error())
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines = make([]string, 0)
	for scanner.Scan() {

		text := scanner.Text()
		lines = append(lines, text)
	}

	log.Println("Done reading file..")
	return lines, nil
}
