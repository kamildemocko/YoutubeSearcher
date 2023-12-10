package types_

type CurrentData struct {
	ID           int
	VideoId      string
	Title        string
	Description  string
	Url          string
	ThumbnailUrl string
	PublishedAt  string
	ChannelID    string
	ChannelTitle string
	ChannelUrl   string
}

type Config struct {
	ApiKey                    string
	YoutubeSearchGetRequest   string
	YoutubeVideoGetRequtest   string
	YoutubeCommentsGetRequest string
	MaxChannelLength          int
	MaxTitleLength            int
	MaxResults                int
}
