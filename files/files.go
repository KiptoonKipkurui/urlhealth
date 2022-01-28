//
// package files contains helper methods to read files
package files

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/kiptoonkipkurui/urlhealth/processing"
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

type ReadMe struct {
	Path string
	Name string

	Links []Link
}

type Link struct {
	Url     string
	Alias   string
	Healthy bool
	Lock    *sync.Mutex
}

//
// crawls file system to get markdown files
//
func Discover(path string, files []ReadMe) ([]ReadMe, error) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isMarkdown(info.Name()) {
			readme := ReadMe{Path: info.Name(), Name: info.Name(), Links: make([]Link, 0)}
			lines, err := Read(path)

			if err == nil {
				for _, line := range lines {
					urls, _ := processing.GetUrls(line)

					for _, url := range urls {
						readme.Links = append(readme.Links, Link{Url: url})
					}
				}
			}
			files = append(files, readme)
		}

		return nil
	})

	return files, nil
}

//
// checks whether the name has a markdown extension
//
func isMarkdown(name string) bool {
	name = strings.ToLower(name)

	if strings.HasSuffix(name, "md") || strings.HasSuffix(name, "markdown") {
		return true
	}
	return false
}
