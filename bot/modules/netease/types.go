package netease

type SongItem struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Artists []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"artists"`
	Album struct {
		ID     int64  `json:"id"`
		Name   string `json:"name"`
		Artist struct {
			Img1v1Url string `json:"img1v1Url"`
		}
	} `json:"album"`
}

type SongDetail struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}

type SongResult struct {
	Code int          `json:"code"`
	Data []SongDetail `json:"data"`
}

type SearchResult struct {
	Code   int `json:"code"`
	Result struct {
		Songs []SongItem `json:"songs"`
		Count int        `json:"songCount"`
	} `json:"result"`
}
