package util

import (
    "fmt"
    "strings"
    "strconv"
    "regexp"
)

// UpdatePathPagination Func
func UpdatePathPagination(path *string, start, size int) {
    
    index := strings.Index(*path, "?");
    if (index < 0) {
        *path = *path + "?"
    }

    addOrChangeStart(path, strconv.Itoa(start))
    addOrChangeSize(path, strconv.Itoa(size))
}

// FindStartPagination add start and size if initual page doen't
// have it and return the initial pagination if already have add to the path
func FindStartPagination(path string) int {
    pattern := regexp.MustCompile(`start=(\d+)`)
    resultSet := pattern.FindStringSubmatch(path)

    if len(resultSet) == 0 {
        return 0;
    }

    newStart, _ := strconv.Atoi(resultSet[1])
    return newStart
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