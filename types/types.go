package types

type Joke struct{
	ID string `json:"id"`
	Joke string `json:"joke"`
	Status int `json:"status"`
}

type Response struct {
	FrameCount int    `json:"frameCount"`
	Error      string `json:"error"`
	Result     []struct {
		Anilist    struct {
			ID int `json:"id"`
			IdMal int `json:"idMal"`
			Title struct {
				Native  string `json:"native"`
				Romaji  string `json:"romaji"`
				English string `json:"english"`
			} `json:"title"`
			IsAdult  bool     `json:"isAdult"`
		}  `json:"anilist"`
		Filename   string  `json:"filename"`
		Episode    int     `json:"episode"`
		From       float64 `json:"from"`
		To         float64 `json:"to"`
		Similarity float64 `json:"similarity"`
		Video      string  `json:"video"`
		Image      string  `json:"image"`
	} `json:"result"`
}