package rest

type Event struct {
		ID string `json:"id"`
		Type string `json:"type"`
		Actor struct{
			ID string `json:"id"`
			Username string `json:"login"`
			Avatar string `json:"avatar_url"`
			Url string `json:"url"`
		} `json:"actor"`
		Repo struct{
			ID string `json:"id"`
			Name string `json:"name"`
			Url string `json:"url"`
		} `json:"repo"`
}

type Events []struct{Event}


type Items []struct{}


type Flag struct {
	User        string
	PerPage     int
}