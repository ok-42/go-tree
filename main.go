package main

import (
	"fmt"
	"os"
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

// Box drawings light vertical
const S_I string = "\u2502"

// Box drawings light vertical and right
const S_K string = "\u251C"

// Box drawings light up and right
const S_L string = "\u2514"

const BLUE string = "\033[1;34m"
const GREEN string = "\033[1;32m"
const RESET_COLOUR string = "\033[0m"

func read(path string, final []bool) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	lenEntries := len(entries)
	for index, entry := range entries {
		isDir := entry.IsDir()
		isFinal := index == lenEntries-1
		var colouredName string
		var name string = entry.Name()
		if isDir {
			colouredName = BLUE + name + RESET_COLOUR
		} else {
			colouredName = name
		}
		var final_ = append(final, isFinal)
		n := len(final_)
		var out string
		for ind, value := range final_ {
			if ind != n-1 {
				if value {
					out += " "
				} else {
					out += S_I
				}
				out += "   "
			} else {
				if value {
					out += S_L
				} else {
					out += S_K
				}
			}
		}
		fmt.Println(out+"\u2500\u2500", colouredName)
		if isDir && !contains(ignorePaths, name) {
			var newPath string
			if path == "." {
				newPath = name
			} else {
				newPath = path + "/" + name
			}
			read(newPath, final_)
		}
	}
}

func main() {
	var path string
	if len(os.Args) == 1 {
		path = "."
	} else {
		path = os.Args[1]
	}
	read(path, []bool{})
}
