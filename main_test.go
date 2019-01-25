package main

import (
	"testing"
	"os"
	"strconv"
	"bytes"
	"io/ioutil"
	"net/http"
	"fmt"
	"net/http/httptest"
)

func TestShouldPanicWithNonArgs(t *testing.T) {

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()
	main()
}
func TestShoundPanicOnlyWithPathInArgs(t *testing.T) {

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
	}()
	
	os.Args = []string{`localhost.com`}
	main()
}
func TestShouldMainPanicErrorWithOnlyStringArgs(t *testing.T) {

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
	}()
	
	os.Args = []string{`exec.go`, `localhost.com`, `error`}
	main()
}

func TestShouldMainPanicWithInacessibleServer(t *testing.T) {

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
	}()
	
	os.Args = []string{`exec.go`, `localhost.com`, strconv.Itoa(100)}
	main()
}

func TestArgs(t *testing.T) {
	
	// path := "https://store.playstation.com/chihiro-api/viewfinder/SA/en/999/STORE-MSF75508-FULLGAMES?start=0&size=5"
	// os.Args = []string{path, strconv.Itoa(5)}

	gamesRepository := [2]string {`{
		"age_limit": 7,
		"name": "Game test",
		"playable_platform": ["PS4"],
		"provider_name": "Activision",
		"star_rating": {
			"score": "4.0",
			"total": "2",
			"count": [{
				"star": 4,
				"count": 2
			}]
		}
	}`, 
	`{
		"age_limit": 18,
		"name": "God of war 3",
		"playable_platform": ["PS3"],
		"provider_name": "Sony",
		"star_rating": {
			"score": "5.0",
			"total": "12",
			"count": [{
				"star": 5,
				"count": 12
			}]
		}
	}`}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		index, e := strconv.ParseInt(req.FormValue("start"), 10, 32)
		if e != nil {
			t.Errorf("Error getting start value from Request: %s", e)
		}

		ioutil.NopCloser(
			bytes.NewBufferString(
				fmt.Sprintf(`{ 
					"links": [%s],
					"size": 1,
					"start": %d,
					"total_results": 2
					}`, 
					gamesRepository[index], 
					index,
				),
			),
		)

		rw.Write([]byte(
			fmt.Sprintf(`{ 
				"links": [%s],
				"size": 1,
				"start": %d,
				"total_results": 2
				}`, 
				gamesRepository[index], 
				index,
			),
		))
	}))
	// Close the server when test finishes
	defer server.Close()

	os.Args = []string{"exec.go", server.URL, strconv.Itoa(1)}

	main()
}