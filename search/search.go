package search

import (
    "encoding/json"
    "github.com/goinggo/tracelog"
    "io/ioutil"
    "net/http"
    "psstore/util"
    "time"
)

//API struct
type API struct {
    Client *http.Client
    URL string
}

// GetGames search and return the result of game searched on playstation store
func getGames(api *API) (Result *ResultSearch, e error) {
    response, err := api.Client.Get(api.URL)
    
    if err == nil {
        data, _ := ioutil.ReadAll(response.Body)
        
        result := ResultSearch{}
        unmarshalErr := json.Unmarshal(data, &result)
        if unmarshalErr != nil {
            return nil, unmarshalErr
        }
        return &result, nil
    }
    return nil, err
}

// Execute the search
func Execute(path string, itemsPerPage int) ([]Game, error) {
    client := &http.Client{}


    start := util.FindStartPagination(path)
    util.UpdatePathPagination(&path, start, itemsPerPage)

    API := API{Client: client, URL: path}

    tracelog.Info("PsStore", "Execute",
        "executing search starting at %d with %d games at path %s", start, itemsPerPage, path)

    return execute(&API)
}

func execute(api *API) ([]Game, error) {
    result, err := getGames(api)

    if err != nil {
        return nil, err
    }

    tracelog.Info("PsStore", "execute","found %d games", result.Total)

    var games = result.Games
    if hasNext(result.Start, result.Size, result.Total) {
        var start, size int = next(result.Start, result.Size, result.Total);
        time.Sleep(6 * time.Second)

        path := api.URL
        util.UpdatePathPagination(&path, start, size)

        tracelog.Info("PsStore", "execute",
            "executing search starting at %d with %d games at path %s", start, size, path)

        nextGames, err := execute(&API{Client: api.Client, URL: path})
        if err != nil {
            return nil, err
        }

        tracelog.Info("PsStore", "execute",
            "append more %d games found", len(nextGames))

        return append(games, nextGames...), nil
    }
    return games, nil
}

func hasNext(start, size, total int) bool {
    if size == 0 || total == 0 {
        return false
    }

    next := start + size

    return next < total;
}

func next(start, size, total int) (int, int) {
    next := start + size

    if next > total {
        return total, 0
    }

    return next, size
}
