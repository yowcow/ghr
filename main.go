package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/yowcow/ghr/xmlparser"
)

var (
	modeRe  = regexp.MustCompile(`\AHEAD(\^*)`)
	Version = "x.x.x"
)

func main() {
	var repo string
	var help bool
	var version bool

	flag.StringVar(&repo, "repo", "", "GitHub repository name, e.g., vim/vim")
	flag.BoolVar(&help, "help", false, "print help")
	flag.BoolVar(&version, "version", false, "print version")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if version {
		fmt.Println(Version)
		os.Exit(0)
	}

	if repo == "" {
		flag.Usage()
		os.Exit(1)
	}

	head := os.Args[len(os.Args)-1]
	if !modeRe.MatchString(head) {
		fmt.Println("specify HEAD, HEAD^, HEAD^^, or similar")
		os.Exit(2)
	}

	result := modeRe.FindStringSubmatch(head)
	v, err := getVersionBeforeHead(repo, len(result[1]))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	fmt.Println(v)
}

func xmlURL(repo string) string {
	return fmt.Sprintf("https://github.com/%s/releases.atom", repo)
}

func getVersionBeforeHead(repo string, n int) (string, error) {
	d, err := fetchXML(xmlURL(repo))
	if err != nil {
		panic(err)
	}

	entries, err := xmlparser.ParseAtom(d)
	if err != nil {
		return "", err
	} else if len(entries) == 0 {
		return "", errors.New("no valid release found")
	}

	return getVersionString(repo, entries[n]), nil
}

func getVersionString(repo string, entry xmlparser.Entry) string {
	var v string
	fmt.Sscanf(entry.Link.URL, "/"+repo+"/releases/tag/v%s", &v)
	return v
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
