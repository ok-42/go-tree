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
	for _, entry := range entries {
		var name string = entry.Name()
		var indent string
		if level == 0 {
			indent = strings.Repeat(" ", level)
		} else {
			indent = "\u2502" + strings.Repeat(" ", level-1)
		}
		fmt.Print(indent)
		fmt.Println("\u251C\u2500\u2500", name)
		if entry.Type().IsDir() && !contains(ignorePaths, name) {
			var newPath string
			if path == "." {
				newPath = name
			} else {
				newPath = path + "/" + name
			}
			read(newPath, level+4)
		}
		// fmt.Println(indent, "\u2502")
	}
}

func main() {
	var path string = os.Args[1]
	read(path, 0)
}
