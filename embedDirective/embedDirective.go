package main

import (
	"embed"
)

//go:embed single_file.txt
var fileString string

//go:embed single_file.txt
var fileByte []byte

//go:embed single_file.txt
//go:embed *.hash
var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("key.hash")
	print(string(content1))
}
