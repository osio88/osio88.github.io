package api

const (
	streamInfoUrl = "https://api.sooplive.com/stream?limit=20&page=1&sort=viewer&languageCodeList="
)

type StreamList struct {
	ChannelID    string `json:"channelId"`
	Nickname     string `json:"nickname"`
	Title        string `json:"title"`
	ProfileURL   string `json:"profileUrl"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Viewer       int    `json:"viewer"`
	LanguageCode string `json:"languageCode"`
}

type StreamInfo struct {
	TotalCnt    int          `json:"totalCnt"`
	CurrentPage int          `json:"currentPage"`
	IsNext      bool         `json:"isNext"`
	StreamList  []StreamList `json:"streamList"`
}

func GetStream(limit, page int, sort, languageCodeList string) *StreamInfo {
	// TODO: Implement this function
	return &StreamInfo{
		TotalCnt:    1,
		CurrentPage: 1,
		IsNext:      true,
		StreamList: []StreamList{
			{
				ChannelID:    "sherlock",
				Nickname:     "Sherlock",
				Title:        "sherlock is broadcasting",
				ProfileURL:   "",
				ThumbnailURL: "",
				Viewer:       0,
				LanguageCode: "en",
			},
			{
				ChannelID:    "endolphin",
				Nickname:     "Endolphin",
				Title:        "endolphin is broadcasting",
				ProfileURL:   "",
				ThumbnailURL: "",
				Viewer:       0,
				LanguageCode: "en",
			},
		},
	}
}
