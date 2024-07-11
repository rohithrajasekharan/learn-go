package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("./directories/test", 0755)
	check(err)

	defer os.RemoveAll("./directories/test")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("./directories/file1")

	err = os.MkdirAll("./directories/sub-directories/child", 0755)
	check(err)

	createEmptyFile("./directories/sub-directories/file2")
	createEmptyFile("./directories/sub-directories/file3")
	createEmptyFile("./directories/sub-directories/child/file4")

	c, err := os.ReadDir("./directories/sub-directories")
	check(err)

	fmt.Println("Listing sub directories")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("./directories/sub-directories/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing directories/sub-directories/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)

}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
