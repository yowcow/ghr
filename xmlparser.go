package main

import (
	"encoding/xml"
)

type Href struct {
	URL string `xml:"href,attr"`
}

type Entry struct {
	Link Href `xml:"link"`
}

type Atom struct {
	Entries []Entry `xml:"entry"`
}

func ParseAtom(data []byte) ([]Entry, error) {
	var atom = new(Atom)

	if err := xml.Unmarshal(data, &atom); err != nil {
		return nil, err
	}

	return atom.Entries, nil
}
