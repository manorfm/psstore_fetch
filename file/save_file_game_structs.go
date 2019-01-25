package file

// Game struct
type Game struct {
    AgeLimit int `json:"age"`
    Name string `json:"name"`
    Plataforms[] string `json:"platform"`
	Genere[] string `json:"genere"`
    ScreenShots[] string `json:"screenshots"`
    DataRelease string `json:"release"`
    ProviderName string `json:"provider"`
    Rating Rating `json:"rating"`
    Price string `json:"price"`
}

// Vote struct
type Vote struct {
    Star int `json:"star"`
    Count int `json:"total"`
}

// Rating struct
type Rating struct {
    Score float32 `json:"score"`
    Total int `json:"total"`
    Votes[] Vote `json:"votes"`
}

// GameData struct
type GameData struct {
    Games[] Game `json:"games"`
    Total int `json:"total"`
}