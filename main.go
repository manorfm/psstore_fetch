package main

import (
    "fmt"
    "log"
    "os"
    "psstore/convert"
    "psstore/file"
    "psstore/search"
    "strconv"
)

var (
    logFile  *os.File
)

func init() {

    logFolder := "./log"
    if _, err := os.Stat(logFolder); os.IsNotExist(err) {
        os.Mkdir(logFolder, os.ModePerm)
    }
    logFile, err := os.OpenFile("./log/psstore.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }

    log.SetOutput(logFile)
}

func main() {
    defer logFile.Close()

    args := os.Args

    if len(args) != 3 {
        exiting("Error: Empty args, I need PS api path and pagination amount by args.", 1)
    }

    path := args[1]
    itemsPerPageStr := args[2]
    
    itemsPerPage, err := strconv.Atoi(itemsPerPageStr)
    
    if err != nil {
        exiting("Error: Pagination amount has to be numerical.", 2)
    }
    
    games, err := search.Execute(path, itemsPerPage)
    
    if err != nil {
        exiting(fmt.Sprintf("Error while execute search to path %s, %v", path, err), 3)
    }

    log.Printf("Fetched a total of %d games", len(games), )

    file.Write(convert.ToFileStructureGames(games))
}

func exiting(message string, code int) {
    log.Println(message)
    os.Exit(code)
}

