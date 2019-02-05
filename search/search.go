package search

import (
    "encoding/json"
    "github.com/schollz/progressbar"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "github.com/manorfm/psstore/util"
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

    log.Printf("executing search starting at %d with %d games at path %s", start, itemsPerPage, path)

    return execute(&API)
}

func execute(api *API) ([]Game, error) {
    result, err := getGames(api)

    if err != nil {
        return nil, err
    }

    bar := progressbar.NewOptions(result.Total)
    bar.RenderBlank()
    bar.Add(result.Size)


    var games = result.Games
    for hasNext(result.Start, result.Size, result.Total) {
        var start, size int = next(result.Start, result.Size, result.Total);

        sleep := random(3, 6)
        time.Sleep(time.Duration(sleep) * time.Second)

        path := api.URL
        util.UpdatePathPagination(&path, start, size + 1)

        log.Printf("Executing search starting at %d and fetching more %d games of %d at path %s", start, size, result.Total, path)

        result, err = getGames(&API{Client: api.Client, URL: path})
        if err != nil {
            return nil, err
        }

        bar.Add(result.Size)
        games = append(games, result.Games...)
    }

    return games, nil
}

func random(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min
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
