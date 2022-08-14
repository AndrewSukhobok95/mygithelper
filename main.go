package main

import (
	"fmt"

	"github.com/AndrewSukhobok95/mygithelper/cmd"
)

var (
	VersionTag string
	Timestamp  string
)

func main() {
	if VersionTag != "" {
		fmt.Printf("Git tag : %s\nBuilt at: %s\n\n", VersionTag, Timestamp)
	}
	cmd.Execute()
}
