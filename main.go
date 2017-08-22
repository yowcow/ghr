package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var XMLURL = "https://github.com/vim/vim/releases.atom"

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Fprintln(os.Stdout, "Usage:", args[0], "latest")
	} else if args[1] == "latest" {
		doLatest()
	} else {
		fmt.Fprintln(os.Stderr, "Unknown mode:", args[1])
	}
}

func doLatest() {
	d, err := fetchXML(XMLURL)

	if err != nil {
		panic(err)
	}

	entries, err := ParseAtom(d)

	if err != nil {
		panic(err)
	} else if len(entries) == 0 {
		panic(errors.New("No release found!  Something should be wrong!"))
	}

	printVersion(entries[0])
}

func printVersion(entry Entry) {
	var version string

	fmt.Sscanf(entry.Link.URL, "/vim/vim/releases/tag/v%s", &version)
	fmt.Println(version)
}

func fetchXML(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Expected status code 200 but got %d", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}
