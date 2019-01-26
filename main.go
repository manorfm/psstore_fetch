package main

import (
    "fmt"
    "os"
    "psstore/convert"
    "psstore/file"
    "psstore/search"
    "strconv"
    "github.com/goinggo/tracelog"
)

func main() {
    fmt.Println("Starting the application...")
    tracelog.StartFile(tracelog.LevelTrace, "./log", 1)

    tracelog.Start(tracelog.LevelTrace)

    argsWithProg := os.Args
    if len(argsWithProg) != 3 {
        exiting("Error: Needed the path and the amount by pagination args.", 1)
    }

    path := argsWithProg[1]
    itemsPerPageStr := argsWithProg[2]
    
    itemsPerPage, err := strconv.Atoi(itemsPerPageStr)
    
    if err != nil {
        exiting("Error: Inform a number to amount for pagination as parameter", 2)
    }
    
    games, err := search.Execute(path, itemsPerPage)
    
    if err != nil {
        exiting(fmt.Sprintf("Error while execute search to path %s, %v", path, err), 3)
    }

    file.Write(convert.ToFileStructureGames(games))

    fmt.Println("End the application...")
}

func exiting(message string, code int) {
    tracelog.Errorf(fmt.Errorf(message), "main", "main", "Hello Error")
    tracelog.Stop()
    os.Exit(1)
}

