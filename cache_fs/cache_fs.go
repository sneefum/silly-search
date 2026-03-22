package cache_fs

import (
	// "fmt"
	"log"
	"errors"
	"strings"
	"os/exec"
	"strconv"

	// "https://github.com/lib/pq"
)

type File struct {
	name string
	creation_time uint64
	size uint64
}

func FileInfo(path string) File {
	cmd := exec.Command("stat", "-c%W,%s:%n", path)
	
	if errors.Is(cmd.Err, exec.ErrDot) {
		log.Panic("Command error!")
		cmd.Err = nil
	}

	output, err := cmd.Output()
	
	if err != nil {
		log.Panic(err)
	}
	
	outputStr := string(output)

	f := File{}
	colonPos := strings.IndexByte(outputStr, ':')
	commaPos := strings.IndexByte(outputStr, ',')

	if colonPos == -1 {
		log.Panic("Colon not found!")
	}
	if commaPos == -1 {
		log.Panic("Comma not found!")
	}

	f.name = outputStr[colonPos+1:]

	f.size, err = strconv.ParseUint(outputStr[commaPos+1:colonPos], 10, 64)
	if err != nil {
		log.Panic("Failed to convert size to uint64!")
	}

	f.creation_time, err = strconv.ParseUint(outputStr[:commaPos], 10, 64)
	if err != nil {
		log.Panic("Failed to convert creation time to uint64!")
	}

	return f
}
