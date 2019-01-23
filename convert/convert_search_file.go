package convert

import (
	"psstore/search"
	"psstore/file"
)

// ToFileStructureGames from search.Game to file.Game
func ToFileStructureGames(games []search.Game) *[]file.Game {
    var fileGames []file.Game
    for _, searchGame := range games {
        fileGames = append(fileGames, file.Game {
            AgeLimit: searchGame.AgeLimit,
            Name: searchGame.Name,
            Plataforms: searchGame.Plataforms,
            Genere: *convertGenere(searchGame.Metadatas),
            ScreenShots: *convertScreeShots(searchGame.Media),
            DataRelease: searchGame.DataRelease,
            ProviderName: searchGame.ProviderName,
            Rating: *convertRating(&searchGame.Rating),
        })
	}
    return &fileGames
}

func convertRating(rating *search.Rating) *file.Rating {
    searchVotes := rating.Votes

    fileRating := file.Rating {
        Score: rating.Score,
        Total: rating.Total,
    }
    
    for _, vote := range searchVotes {
        fileRating.Votes = append(fileRating.Votes, file.Vote {
            Star: vote.Star,
            Count: vote.Count,
        })
    }

    return &fileRating
}
func convertGenere(metadata search.Metadata) *[]string {
    searchGenere := metadata.Genere

    var generes []string
    for _, genere := range searchGenere.Values {
        generes = append(generes, genere)
    }

    return &generes
}

func convertScreeShots(media search.Media) *[]string {
    searchScreenShots := media.ScreenShots

    var screenShots []string
    for _, screenShot := range searchScreenShots {
        screenShots = append(screenShots, screenShot.URL)
    }

    return &screenShots
}