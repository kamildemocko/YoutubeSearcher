package types_

type TargetSearchYoutube struct {
	Kind          string
	Etag          string
	NextPageToken string
	RegionCode    string
	PageInfo      struct {
		TotalResults  int
		ResultPerPage int
	}
	Items []ItemSearchYoutube
}

type ItemSearchYoutube struct {
	Kind string
	Etag string
	Id   struct {
		Kind    string
		VideoId string
	}
	Snippet struct {
		PublishedAt string
		ChannelId   string
		Title       string
		Description string
		Thumbnails  struct {
			Default struct {
				Url    string
				Width  int
				Height int
			}
			Medium struct {
				Url    string
				Width  int
				Height int
			}
			High struct {
				Url    string
				Width  int
				Height int
			}
		}
		ChannelTitle         string
		LiveBroadcastContent string
		PublishTime          string
	}
}

type TargetVideoYoutubeStatistics struct {
	Kind     string
	Etag     string
	Items    []ItemVideoYoutubeStatistics
	PageInfo struct {
		TotalResults   int
		ResultsPerPage int
	}
}

type ItemVideoYoutubeStatistics struct {
	Kind       string
	Etag       string
	Id         string
	Statistics struct {
		ViewCount     string
		LikeCount     string
		FavoriteCount string
		CommentCount  string
	}
}

type TargetVideoYoutubeSnippet struct {
	Kind     string
	Etag     string
	Items    []ItemVideoYoutubeSnippet
	PageInfo struct {
		TotalResults   int
		ResultsPerPage int
	}
}

type ItemVideoYoutubeSnippet struct {
	Kind    string
	Etag    string
	Id      string
	Snippet struct {
		PublishedAt string
		ChannelId   string
		Title       string
		Description string
		Thumbnails  struct {
			Default struct {
				Url    string
				Width  int
				Height int
			}
			Medium struct {
				Url    string
				Width  int
				Height int
			}
			High struct {
				Url    string
				Width  int
				Height int
			}
			Standard interface{}
			Maxres   interface{}
		}
		ChannelTitle         string
		LiveBroadcastContent string
		DefaultLanguage      string
		Localized            struct {
			Title       string
			Description string
		}
		DefaultAudioLanguage string
	}
}
