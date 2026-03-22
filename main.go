package main

import (
	"fmt"

	"silly-search/cache_fs"
)

func main() {
	path := "main.go"

	testString := cache_fs.FileInfo(path)
	fmt.Println(testString)
}
