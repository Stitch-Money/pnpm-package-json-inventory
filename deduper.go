package main

import "fmt"

type Deduper map[string]struct{}

func NewDeduper() (d Deduper) {
	return make(Deduper)
}

func (d Deduper) AlreadyParsed(name, version string) bool {
	id := fmt.Sprintf("%s@%s", name, version)

	if _, ok := d[id]; ok {
		return true
	}

	d[id] = struct{}{}

	return false
}
