package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/yowcow/ghr/xmlparser"
)

var (
	modeRe  = regexp.MustCompile(`\AHEAD(\^*)`)
	Version = "x.x.x"
)
var (
	repo    string
	help    bool
	verbose bool
	version bool
)

func init() {
	flag.StringVar(&repo, "repo", "", "GitHub repository name, e.g., vim/vim")
	flag.BoolVar(&help, "help", false, "print help")
	flag.BoolVar(&verbose, "verbose", false, "verbose")
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
}

func main() {
	url := xmlURL(repo)
	d, err := fetchXML(url)
	if err != nil {
		log.Fatalln("failed fetching XML:", err)
	}
	if verbose {
		fmt.Println(string(d))
	}

	entries, err := xmlparser.ParseAtom(d)
	if err != nil {
		log.Fatalln("failed parsing XML:", err)
	} else if len(entries) == 0 {
		log.Fatalln("no valid release found for", url)
	}
	if verbose {
		fmt.Println(entries)
	}

	head := os.Args[len(os.Args)-1]
	if modeRe.MatchString(head) {
		matched := modeRe.FindStringSubmatch(head)
		idx := len(matched[1])
		if len(entries) < idx+1 {
			log.Fatalln("no version for specified", head)
			fmt.Println(getVersionString(repo, entries[idx]))
		}
		return
	}

	for _, entry := range entries {
		fmt.Println(getVersionString(repo, entry))
	}
}

func xmlURL(repo string) string {
	return fmt.Sprintf("https://github.com/%s/releases.atom", repo)
}

func getVersionString(repo string, entry xmlparser.Entry) string {
	var v string
	fmt.Sscanf(entry.Link.URL, "https://github.com/"+repo+"/releases/tag/%s", &v)
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
