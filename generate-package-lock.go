package main

import (
	"fmt"
	"log"
)

type NpmList struct {
	Name            string                   `json:"name"`
	Version         string                   `json:"version"`
	LockfileVersion int                      `json:"lockfileVersion"`
	Dependencies    map[string]NpmDependency `json:"dependencies,omitempty"`
}

type NpmDependency struct {
	Version      string                   `json:"version"`
	Resolved     string                   `json:"resolved,omitempty"`
	Dev          bool                     `json:"dev,omitempty"`
	Requires     map[string]string        `json:"requires,omitempty"`
	Dependencies map[string]NpmDependency `json:"dependencies,omitempty"`
}

func GenerateNpmList(in []PnpmList) (out *NpmList, err error) {
	deduper := NewDeduper()

	if len(in) == 0 {
		return nil, fmt.Errorf("expected one pnpm list")
	}
	if len(in) > 1 {
		log.Println("More than one pnpm list received - using first one")
	}

	l := in[0]
	out = &NpmList{
		Name:            l.Name,
		Version:         l.Version,
		LockfileVersion: 2,
		Dependencies:    make(map[string]NpmDependency),
	}

	// Preload deduper with top-level dependencies so they take precedence
	for name, dep := range l.Dependencies {
		deduper.AlreadyParsed(name, dep.Version)
	}
	for name, dep := range l.DevDependencies {
		deduper.AlreadyParsed(name, dep.Version)
	}

	for name, dep := range l.Dependencies {
		out.Dependencies[name] = translateNpmDependency(&dep, false, deduper)
	}

	for name, dep := range l.DevDependencies {
		out.Dependencies[name] = translateNpmDependency(&dep, true, deduper)
	}

	return
}

func translateNpmDependency(in *PnpmDependency, dev bool, deduper Deduper) (out NpmDependency) {
	out.Version = in.Version
	out.Resolved = in.Resolved
	out.Dev = dev
	out.Requires = make(map[string]string)
	out.Dependencies = make(map[string]NpmDependency)

	for name, dep := range in.Dependencies {
		if deduper.AlreadyParsed(name, dep.Version) {
			out.Requires[name] = dep.Version
		} else {
			out.Dependencies[name] = translateNpmDependency(&dep, dev, deduper)
		}
	}

	return
}
