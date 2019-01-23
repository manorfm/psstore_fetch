package util

import (
    //"fmt"
    "strconv"
    "regexp"
)

// ReplacePathPagination Func
func ReplacePathPagination(path string, start, size int) string {
    regexStart := regexp.MustCompile("(start=)\\d+")
	regexSize := regexp.MustCompile("(size=)\\d+")
    
    startReplaced := regexStart.ReplaceAllString(path, "${1}" + strconv.Itoa(start))
    newPath := regexSize.ReplaceAllString(startReplaced, "${1}" + strconv.Itoa(size))
    
    return newPath
}