package convert

import (
	"psstore/search"

	"github.com/stretchr/testify/assert"
	"testing"
)

func getGames() *[]search.Game {
	return &[]search.Game {
		{
			Name: "God of War 4",
			Media: search.Media {
				ScreenShots: []search.ScreenShot {
					{
						URL: "localhost/test1",
					},
					{
						URL: "localhost/test2",
					},
				},
			},
			Plataforms: []string {"PS4"},
			Sku: search.Sku {
				Price: "56.00",
			},
		},
		{
			Name: "Fifa 2019",
			Metadatas: search.Metadata {
				Genere: search.Genere {
					Values: []string {"Genre 1", "Genre 2"},
				},
			},
			Plataforms: []string {"PS4", "XBox One"},
			Rating: search.Rating {
				Score: 4,
				Total: 2,
				Votes: []search.Vote {
					{
						Star: 4,
						Count: 2,
					},
				},
			},
		},
	}
}

func TestConvertingRatingVotes(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game := fileGames[1]
	rating := game.Rating


	assert.Equal(t, float32(4), rating.Score, "The score of rating from file does not match")
	assert.Equal(t, 2, rating.Total, "The score of rating from file does not match")
	assert.Equal(t, 1, len(rating.Votes), "votes counts does not match")

	vote := rating.Votes[0]
	assert.Equal(t, 4, vote.Star, "The start of rating from file does not match")
	assert.Equal(t, 2, vote.Count, "The count of rating from file does not match")
}

func TestConvertingGamesAndTestLength(t *testing.T) {
	searchGames := *getGames()
	fileGames := ToFileStructureGames(searchGames)

	assert.Equal(t, 2, len(*fileGames), "written json does not match")
}
func TestPlataforms(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game0 := fileGames[0]

	assert.Equal(t, 1, len(game0.Plataforms), "written json does not match")
	assert.Equal(t, "PS4", game0.Plataforms[0], "Plataform PS4 does not match")


	game1 := fileGames[1]
	assert.Equal(t, 2, len(game1.Plataforms), "written json does not match")
}

func TestConvertMetadataToGenre(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game := fileGames[1]
	assert.Equal(t, "Genre 1", game.Genere[0], "genre 1 does not match")
	assert.Equal(t, "Genre 2", game.Genere[1], "genre 2 does not match")
}

func TestConvertMediaToScreenShots(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game := fileGames[0]
	assert.Equal(t, "localhost/test1", game.ScreenShots[0], "screenshot 1 does not match")
	assert.Equal(t, "localhost/test2", game.ScreenShots[1], "screenshot 2 does not match")
}

func TestPrice(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game := fileGames[0]
	assert.Equal(t, "56.00", game.Price, "the game price does not match")
}
