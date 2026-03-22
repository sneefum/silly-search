package cache_fs

import (
	// "fmt"
	"log"
	"errors"
	"os/exec"

	// "https://github.com/lib/pq"
)

func FileInfo(path string) string {
	cmd := exec.Command("stat", "-c%W,%s,%n", path)
	
	if errors.Is(cmd.Err, exec.ErrDot) {
		log.Panic("Command error!")
		cmd.Err = nil
	}

	output, err := cmd.Output()
	
	if err != nil {
		log.Panic(err)
	}
	
	outputStr := string(output)
	return outputStr
}
