package convert

import (
	"testing"
	"psstore/search"
)

func getGames() *[]search.Game {
	return &[]search.Game {
		search.Game {
			Name: "God of War 4",
			Media: search.Media {
				ScreenShots: []search.ScreenShot {
					search.ScreenShot {
						URL: "localhost/test1",
					},
					search.ScreenShot {
						URL: "localhost/test2",
					},
				},
			},
			Plataforms: []string {"PS4"},
			Sku: search.Sku {
				Price: "56.00",
			},
		},
		search.Game {
			Name: "Fifa 2019",
			Metadatas: search.Metadata {
				Genere: search.Genere {
					Values: []string {"Genere 1", "Genere 2"},
				},
			},
			Plataforms: []string {"PS4", "XBox One"},
			Rating: search.Rating {
				Score: 4,
				Total: 2,
				Votes: []search.Vote {
					search.Vote {
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


	if rating.Score != 4 {
		t.Errorf("The score of rating from file does not match was %f, want: %d.", rating.Score, 4)
	}
	if rating.Total != 2 {
		t.Errorf("The score of rating from file does not match was %d, want: %d.", rating.Total, 2)
	}
	
	if len(rating.Votes) != 1 {
		t.Errorf("votes counts does not match was %d, want: %d.", len(rating.Votes), 1)
	}
	
	vote := rating.Votes[0]
	if vote.Star != 4 {
		t.Errorf("The start of rating from file does not match was %d, want: %d.", vote.Star, 4)
	}
	if vote.Count != 2 {
		t.Errorf("The count of rating from file does not match was %d, want: %d.", vote.Count, 2)
	}
}
func TestConvertingGamesAndTestLength(t *testing.T) {
	searchGames := *getGames()
	fileGames := ToFileStructureGames(searchGames)

	if len(*fileGames) != 2 {
		t.Errorf("written json does not match")
	}
}
func TestPlataforms(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game0 := fileGames[0]
	if len(game0.Plataforms) != 1 {
		t.Errorf("written json does not match")
	}
	
	if game0.Plataforms[0] != "PS4" {
		t.Errorf("Plataform PS4 json does not match")
	}


	game1 := fileGames[1]
	if len(game1.Plataforms) != 2 {
		t.Errorf("written json does not match")
	}
}

func TestConvertMetadataToGenere(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game := fileGames[1]
	if game.Genere[0] != "Genere 1" {
		t.Errorf("genere does not match")
	}
	if game.Genere[1] != "Genere 2" {
		t.Errorf("genere does not match")
	}
}

func TestConvertMediaToScreenShots(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game := fileGames[0]
	if game.ScreenShots[0] != "localhost/test1" {
		t.Errorf("media screenshots not match")
	}
	if game.ScreenShots[1] != "localhost/test2" {
		t.Errorf("media screenshots does not match")
	}
}

func TestPrice(t *testing.T) {
	searchGames := *getGames()
	fileGames := *ToFileStructureGames(searchGames)

	game := fileGames[0]
	if game.Price != "56.00" {
		t.Errorf("the price does not match, receive %s, want %s", game.Price, "56.00")
	}
}
