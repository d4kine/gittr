// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	root := "C:/Projects/private/gittr"

	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() {
			fmt.Println(path)
		}

		return nil
	})

}
