package util

import (
    "fmt"
    "strings"
    "strconv"
    "regexp"
)

// ReplacePathPagination Func
func ReplacePathPagination(path string, start, size int) string {
    
    index := strings.Index(path, "?");
    if (index < 0) {
        path = path + "?"
    }

    addOrChangeStart(&path, strconv.Itoa(start))
    addOrChangeSize(&path, strconv.Itoa(size))
    
    return path
}

func addOrChangeStart(path *string, start string) {
    regexStart := regexp.MustCompile("(start=)\\d+")
    startLoc := regexStart.FindIndex([]byte(*path))

    if len(startLoc) > 0 {
        *path = regexStart.ReplaceAllString(*path, "${1}" + start)
    } else if strings.LastIndex(*path, "size=") > 0 {
        *path = fmt.Sprintf("%s&start=%s", *path, start)
    } else {
        *path = fmt.Sprintf("%sstart=%s", *path, start)
    }
}

func addOrChangeSize(path *string, size string) {
    regexSize := regexp.MustCompile("(size=)\\d+")
    sizeLoc := regexSize.FindIndex([]byte(*path))

    if len(sizeLoc) > 0 {
        *path = regexSize.ReplaceAllString(*path, "${1}" + size)
    } else {
        *path = fmt.Sprintf("%s&size=%s", *path, size)
    }
}