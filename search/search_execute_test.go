package search

import (
	"bytes"
	"github.com/goinggo/tracelog"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"fmt"
	"strconv"
	"net/http/httptest"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestDoStuffWithRoundTripper(t *testing.T) {

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

	var countRequestsDid = 0
	client := NewTestClient(func(req *http.Request) *http.Response {

		countRequestsDid = countRequestsDid + 1

		index, e := strconv.ParseInt(req.FormValue("start"), 10, 32)
		if e != nil {
			t.Errorf("Error getting start value from Request: %s", e)
		}

		return &http.Response {
			StatusCode: 200,
			Body: ioutil.NopCloser(
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
			),
			Header: make(http.Header),
		}
	})

	api := API{client, "http://localhost/test?start=0&size=1"}
	games, _ := execute(&api)
	
	length := len(games)
	if length != 2 {
		t.Errorf("The count of games should was %d, want: %d.", length, 2)
	}
	if countRequestsDid != 2 {
		t.Errorf("The count of requests should was %d, want: %d.", countRequestsDid, 2)
	}	
}

func TestRequestErrorExecuteFunc(t *testing.T) {	
	_, e := Execute("http://localhost/something/testing", 10)
	assertNotNull(e, "should return a error", t)
}

func TestUnMarshalingErrorExecuteFunc(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		rw.Write([]byte(`{
			"links": [{
				"age_limit": "error7"
			}]
		}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	_, e := Execute(server.URL, 20)
	
	assertNotNull(e, "should return a error", t)
}

func TestTheNextIterationUnMarshalingErrorExecuteFunc(t *testing.T) {

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
		"age_limit": provokingError18,
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
			),
		)
	}))
	// Close the server when test finishes
	defer server.Close()

	_, e := Execute(server.URL, 1)
	
	assertNotNull(e, "should return a error", t)
}

func TestMain(m *testing.M) {
	tracelog.Start(tracelog.LevelTrace)
	code := m.Run()
	tracelog.Stop()
	os.Exit(code)
}