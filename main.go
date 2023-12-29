package main

import (
	"fmt"
	"os"
	"strings"
)

// https://stackoverflow.com/a/10485970
func contains[anyType int | string](s []anyType, e anyType) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var ignorePaths = []string{
	"__pycache__",
	".git",
	".idea",
	".ipynb_checkpoints",
	"venv",
}

func read(path string, level int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	lenEntries := len(entries)
	for index, entry := range entries {
		isFinal := index == lenEntries-1
		var name string = entry.Name()
		var indent string = strings.Repeat("\u2502   ", level)
		fmt.Print(indent)
		var finalOrContinue string
		if isFinal {
			finalOrContinue = "\u2514"
		} else {
			finalOrContinue = "\u251C"
		}
		fmt.Println(finalOrContinue+"\u2500\u2500", name)
		if entry.Type().IsDir() && !contains(ignorePaths, name) {
			var newPath string
			if path == "." {
				newPath = name
			} else {
				newPath = path + "/" + name
			}
			read(newPath, level+1)
		}
	}
}

func main() {
	var path string = os.Args[1]
	read(path, 0)
}
