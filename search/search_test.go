package search

import (
	"net/http"

	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
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

	assert.Nil(t, e, "We got some problem, search has ended with some error: \"%s\"")
	assert.Equal(t, 10, result.Size, "Size of search result was incorrect")
	assert.Equal(t, 0, result.Start, "Start of search result was incorrect")
	assert.Equal(t, 40, result.Total, "Total of search result was incorrect")
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

	assert.Nil(t, e, "We got some problem, search has ended with some error: \"%s\"", e)
	assert.Equal(t, 1, len(result.Games), "Quantity of games result from search was incorrect")
	
	game := result.Games[0]

	assert.Equal(t, 7, game.AgeLimit,"Age limit of games result from search was incorrect")
	assert.Equal(t, "Game test", game.Name,"Name of game of games result from search was incorrect")
	assert.Equal(t, 1, len(game.Plataforms),"Plataform of games result from search was incorrect")

	plataform := game.Plataforms[0]
	assert.Equal(t, "PS4", plataform, "Plataform of game of games result from search was incorrect")
	assert.Equal(t, "Activision", game.ProviderName, "Provider of game of games result from search was incorrect")

	rating := game.Rating

	assert.NotNil(t, &rating, "Error by: Rating of games result not exist")
	assert.Equal(t, float32(4.0), rating.Score, "Rating score of rating of games result from search was incorrect")
	assert.Equal(t, 2, rating.Total, "Total votes of rating of games result from search was incorrect")

	votes := rating.Votes

	assert.Equal(t, 1, len(votes), "Votes rating of games result from search was incorrect")

	vote := votes[0]
	assert.Equal(t, 4, vote.Star, "Total start of rating of games result from search was incorrect")
	assert.Equal(t, 2, vote.Count, "Total count of rating of games result from search was incorrect")
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
	
	assert.NotNil(t, e, "should return a error")
}

func TestRequestError(t *testing.T) {
	client := &http.Client{}
	API := API{Client: client, URL: "http://localhost/something/testing"}
	
	_, e := getGames(&API)
	assert.NotNil(t, e, "should return a error")
}

func TestNext(t *testing.T) {
	next, _ := next(0, 10, 5242);
	
	assert.Equal(t, 10, next, "Next was incorrect")
}
func TestNextAboveTotal(t *testing.T) {
	next, size := next(5240, 10, 5242)
	
	assert.Equal(t, 5242, next, "Next was incorrect")
	assert.Equal(t, 0, size, "Next was incorrect")
}

func TestNotHasNext(t *testing.T) {
	assert.False(t, hasNext(5300, 10, 5242), "HasNext was incorrect")
}

func TestNotHasNextTotalZero(t *testing.T) {
	assert.False(t, hasNext(0, 10, 0), "HasNext can't have total equals 0")
}

func TestNotHasNextSizeZero(t *testing.T) {
	assert.False(t, hasNext(0, 0, 10), "HasNext can't have size equals 0")
}

func TestHasNext(t *testing.T) {
	assert.True(t, hasNext(0, 10, 5242), "HasNext was incorrect")
}