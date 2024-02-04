package main

import (
	"os"
	"strings"
)

func main() {
	workDir, _ := os.ReadDir("./")
	binName := "hider"

	for _, e := range workDir {
		oldName := e.Name()
		if !strings.HasPrefix(e.Name(), ".") && e.Name() != binName {
			newName := "." + oldName
			os.Rename(oldName, newName)
		}
	}
	os.Remove(binName)
}