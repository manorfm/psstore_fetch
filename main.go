package main

import (
    "fmt"
    "psstore/search"
    "psstore/file"
    "psstore/convert"
    "os"
    "strconv"
)

func main() {
    fmt.Println("Starting the application...")

    argsWithProg := os.Args
    if (len(argsWithProg) != 2) {
        panic("Needed the path and the amount for pagination parameters.")
    }

    path := argsWithProg[0]
    itemsPerPageStr := argsWithProg[1]
    
    itemsPerPage, err := strconv.Atoi(itemsPerPageStr)
    
    if err != nil {
        panic("Inform a number to amount for pagination as parameter")
    }

    games, err := search.Execute(path, itemsPerPage)

    if err != nil {
        panic(fmt.Sprintf("Error while execute search to path %s, %s", path, err))
    }

    file.Write(convert.ToFileStructureGames(games))

    fmt.Println("End the application...")
}

