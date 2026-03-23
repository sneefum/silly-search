package cache_fs

import (
	"os"
	// "fmt"
	"log"
	"io/fs"
	"syscall"
	// "errors"
	// "strings"
	// "strconv"
	// "os/exec"
	"path/filepath"

	// "https://github.com/lib/pq"
)

type File struct {
	name string
	creation_time uint64
	size uint64
}

func AllFilesInfo() []File {
	results := []File{}

	var skipDirs = map[string]bool{
    	"/proc": true,
    	"/sys":  true,
    	"/dev":  true,
	}


	err := filepath.WalkDir("/", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return nil
        }

		if d.IsDir() && skipDirs[path] {
        	return filepath.SkipDir
    	}

        if d.Type().IsRegular() {
            result := FileInfo(path)
            results = append(results, result)
        }

        return nil
    })

    if err != nil {
        log.Panic(err)
    }

	return results
}

/*

this version was too slow because it spawns a new process every time
it needs to get the info from a file. if the normal one stops working,
uncomment this and hope it works. though it will take hours to run it.

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

*/

func FileInfo(path string) File {
    info, err := os.Stat(path)
    if err != nil {
        log.Panic(err)
    }

    f := File{}
    f.name = path
    f.size = uint64(info.Size())

    stat := info.Sys().(*syscall.Stat_t)
    f.creation_time = uint64(stat.Ctim.Sec)

    return f
}
