package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "C:/Projects/private/gittr"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}


		if d.IsDir() && !strings.Contains(path, ".git/") && strings.Contains(path, ".git"){
			filename := filepath.Join(root, d.Name())
			fmt.Println(filename)
		}

		return nil
	})

}

func usage() {

}