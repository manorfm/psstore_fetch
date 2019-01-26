package search

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchTotal(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"size": 10, "start": 0, "total_results": 40}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{server.Client(), server.URL}
	result, e := getGames(&api)

	assertNull(e, "We got some problem, search has ended with some error: \"%s\"", t)
	assert(result.Size, 10, "Size of search result was incorrect", t)
	assert(result.Start, 0, "Start of search result was incorrect", t)
	assert(result.Total, 40, "Total of search result was incorrect", t)
}

func TestSearchGame(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		rw.Write([]byte(`{
			"links": [{
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
			}]
		}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := API{server.Client(), server.URL}
	result, e := getGames(&api)

	assertNull(e, "We got some problem, search has ended with some error: \"%s\"", t)
	assert(len(result.Games), 1, "Quantity of games result from search was incorrect", t)
	
	game := result.Games[0]

	assert(game.AgeLimit, 7, "Age limit of games result from search was incorrect", t)
	assertStr(game.Name, "Game test", "Name of game of games result from search was incorrect", t)
	assert(len(game.Plataforms), 1, "Plataform of games result from search was incorrect", t)

	plataform := game.Plataforms[0]
	assertStr(plataform, "PS4", "Plataform of game of games result from search was incorrect", t)
	assertStr(game.ProviderName , "Activision", "Provider of game of games result from search was incorrect", t)

	rating := game.Rating

	assertNotNull(&rating, "Error by: Rating of games result not exist", t)
	assertFloat(rating.Score, 4.0, "Rating score of rating of games result from search was incorrect", t)
	assert(rating.Total, 2, "Total votes of rating of games result from search was incorrect", t)

	votes := rating.Votes

	assert(len(votes), 1, "Votes rating of games result from search was incorrect", t)

	vote := votes[0]
	assert(vote.Star, 4, "Total start of rating of games result from search was incorrect", t)
	assert(vote.Count, 2, "Total count of rating of games result from search was incorrect", t)
}

func TestUnMarshalingError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		rw.Write([]byte(`{
			"links": [{
				"age_limit": "error7"
			}]
		}`))
	}))
	// Close the server when test finishes
	defer server.Close()

	api := API{server.Client(), server.URL}
	_, e := getGames(&api)
	
	assertNotNull(e, "should return a error", t)
}

func TestRequestError(t *testing.T) {
	client := &http.Client{}
	API := API{Client: client, URL: "http://localhost/something/testing"}
	
	_, e := getGames(&API)
	assertNotNull(e, "should return a error", t)
}

func TestNext(t *testing.T) {
	next, _ := next(0, 10, 5242);
	
	assert(next, 10, "Next was incorrect", t)
}
func TestNextAboveTotal(t *testing.T) {
	next, size := next(5240, 10, 5242)
	
	assert(next, 5242, "Next was incorrect", t)
	assert(size, 0, "Next was incorrect", t)
}

func TestNotHasNext(t *testing.T) {
	assertFalse(hasNext(5300, 10, 5242), "HasNext was incorrect", t)
}

func TestNotHasNextTotalZero(t *testing.T) {
	assertFalse(hasNext(0, 10, 0), "HasNext can't have total equals 0", t)
}

func TestNotHasNextSizeZero(t *testing.T) {
	assertFalse(hasNext(0, 0, 10), "HasNext can't have size equals 0", t)
}

func TestHasNext(t *testing.T) {
	assertTrue(hasNext(0, 10, 5242), "HasNext was incorrect", t)
}

type Object interface {

}

func assertNotNull(value Object, message string, t *testing.T) {
	if value == nil {
		t.Errorf(message, value)
	}
}

func assertNull(value Object, message string, t *testing.T) {
	if value != nil {
		t.Errorf(message, value)
	}
}

func assert(value, compare int, message string, t *testing.T) {
	if value != compare {
		t.Errorf("%s, got: %d, want: %d.", message, value, compare)
	}
}

func assertFloat(value, compare float32, message string, t *testing.T) {
	if value != compare {
		t.Errorf("%s, got: %f, want: %f.", message, value, compare)
	}
}
func assertStr(value, compare string, message string, t *testing.T) {
	if value != compare {
		t.Errorf("%s, got: %s, want: %s.", message, value, compare)
	}
}

func assertFalse(value bool, message string, t *testing.T) {
	if value {
		t.Errorf(message)
	}
}

func assertTrue(value bool, message string, t *testing.T) {
	assertFalse(!value, message, t)
}