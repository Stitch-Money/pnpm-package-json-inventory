package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	pnpmList, err := GetPnpmList()
	if err != nil {
		log.Fatal(err)
	}

	npmList, err := GenerateNpmList(pnpmList)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(npmList)
}
