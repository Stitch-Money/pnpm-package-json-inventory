package main

import (
	"encoding/json"
	"io"
	"os/exec"
)

type PnpmList struct {
	Name            string
	Version         string
	Path            string
	Private         bool
	Dependencies    map[string]PnpmDependency
	DevDependencies map[string]PnpmDependency
}

type PnpmDependency struct {
	From         string
	Version      string
	Resolved     string
	Dependencies map[string]PnpmDependency
}

func runPnpmList() (io.Reader, error) {
	cmd := exec.Command("pnpm", "list", "--json", "--depth=9999")
	rd, errStdout := cmd.StdoutPipe()

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return rd, errStdout
}

func GetPnpmList() (list []PnpmList, err error) {
	rd, err := runPnpmList()
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(rd)
	err = dec.Decode(&list)

	if err == io.EOF {
		err = nil
	}

	return
}
