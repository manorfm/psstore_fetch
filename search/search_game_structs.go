package search

// Game struct
type Game struct {
    AgeLimit int `json:"age_limit"`
    Name string `json:"name"`
    Plataforms[] string `json:"playable_platform"`
    Metadatas Metadata `json:"metadata"`
    Media Media `json:"mediaList"`
    DataRelease string `json:"release_date"`
    ProviderName string `json:"provider_name"`
    Rating Rating `json:"star_rating"`
    Sku Sku `json:"default_sku"`
}

// Sku informations
type Sku struct {
    Price string `json:"display_price"`
    Type string `json:"type"`
    Available string `json:"playability_date"`
}

// Metadata struct
type Metadata struct {
    Genere Genere `json:"genre"`
}

// Genere genere
type Genere struct {
    Values[] string `json:"values"`
}

//Media struct
type Media struct {
    ScreenShots[] ScreenShot `json:"screenshots"`
}

// ScreenShot struct
type ScreenShot struct {
    URL string `json:"url"`
}

// Vote struct
type Vote struct {
    Star int `json:"star"`
    Count int `json:"count"`
}

// Rating struct
type Rating struct {
    Score float32 `json:"score,string"`
    Total int `json:"total,string"`
    Votes[] Vote `json:"count"`
}

// ResultSearch struct
type ResultSearch struct {
    Games[] Game `json:"links"`
    Start int `json:"start"`
    Size int `json:"size"`
    Total int `json:"total_results"`
}