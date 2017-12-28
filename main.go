package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/yowcow/vimver/xmlparser"
)

var xmlURL = "https://github.com/vim/vim/releases.atom"
var modeRe = regexp.MustCompile(`\AHEAD(\^*)`)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Fprintln(os.Stdout, "Usage:", args[0], "HEAD^^^^")
	} else if result := modeRe.FindStringSubmatch(args[1]); result != nil {
		printVersionBeforeHead(len(result[1]))
	} else {
		fmt.Fprintln(os.Stderr, "Unknown mode:", args[1])
	}
}

func printVersionBeforeHead(n int) {
	d, err := fetchXML(xmlURL)
	if err != nil {
		panic(err)
	}

	entries, err := xmlparser.ParseAtom(d)
	if err != nil {
		panic(err)
	} else if len(entries) == 0 {
		panic(errors.New("no valid release found"))
	}

	printVersion(entries[n])
}

func printVersion(entry xmlparser.Entry) {
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
		return nil, fmt.Errorf("expected status code 200 but got %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
