package main

import (
	"fmt"

	"silly-search/cache_fs"
)

func main() {
	// path := "main.go"

	test := cache_fs.AllFilesInfo()
	fmt.Println("%v", len(test))
}
