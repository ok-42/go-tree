package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

// https://stackoverflow.com/a/37563128
func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
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

// Modified copy of os.ReadDir
func osReadDir(name string, sort_ bool) ([]os.DirEntry, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dirs, err := f.ReadDir(-1)
	if sort_ {
		sort.Slice(dirs, func(i, j int) bool {
			if dirs[i].IsDir() && !dirs[j].IsDir() {
				return true
			} else if !dirs[i].IsDir() && dirs[j].IsDir() {
				return false
			} else {
				return dirs[i].Name() < dirs[j].Name()
			}
		})
	}

	// If .gitignore file is found, ignore its content
	var localPaths []string
	if slices.ContainsFunc[[]os.DirEntry, os.DirEntry](
		dirs, func(de os.DirEntry) bool { return de.Name() == ".gitignore" },
	) {
		gitIgnore, err := os.Open(name + "/.gitignore")
		if err != nil {
			return nil, err
		}
		defer gitIgnore.Close()
		var newPaths []string
		scanner := bufio.NewScanner(gitIgnore)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 0 && !strings.HasPrefix(line, "#") {
				if strings.HasPrefix(line, "/") {
					line = line[1:]
				}
				if strings.HasSuffix(line, "/") {
					line = line[:len(line)-1]
				}
				newPaths = append(newPaths, line)
			}
		}
		localPaths = append(ignorePaths, newPaths...)
	} else {
		localPaths = ignorePaths
	}
	dirs = filter[os.DirEntry](dirs, func(de os.DirEntry) bool {
		return !slices.Contains[[]string, string](localPaths, de.Name())
	})
	return dirs, err
}

func read(path string, final []bool) {
	entries, err := osReadDir(path, true)
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
		var sb strings.Builder
		for _, value := range final {
			if value {
				sb.WriteString("    ")
			} else {
				sb.WriteString(S_I + "   ")
			}
		}
		if isFinal {
			sb.WriteString(S_L)
		} else {
			sb.WriteString(S_K)
		}
		fmt.Println(sb.String()+"\u2500\u2500", colouredName)
		if isDir {
			var newPath string
			if path == "." {
				newPath = name
			} else {
				newPath = path + "/" + name
			}
			read(newPath, append(final, isFinal))
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
