package req

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kamildemocko/src/types_"
)

func SeachYoutube(query string, maxResults int, apiKey string, youtubeSearchGetRequest string) []types_.ItemSearchYoutube {
	response, err := http.Get(buildSearchYoutubeUri(query, maxResults, apiKey, youtubeSearchGetRequest))
	if err != nil {
		panic(err)
	}

	var target types_.TargetSearchYoutube
	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	return target.Items
}

func buildSearchYoutubeUri(query string, maxResults int, apiKey string, youtubeSearchGetRequest string) string {
	params := url.Values{
		"key":        {apiKey},
		"part":       {"snippet"},
		"q":          {query},
		"maxResults": {fmt.Sprintf("%d", maxResults)},
	}

	return youtubeSearchGetRequest + "?" + params.Encode()
}

func GetVideoInfo(videoId string, part string, apiKey string, youtubeVideoGetRequest string) []types_.ItemVideoYoutubeStatistics {
	response, err := http.Get(buildVideoYoutubeUri(videoId, part, apiKey, youtubeVideoGetRequest))
	if err != nil {
		panic(err)
	}

	var target types_.TargetVideoYoutubeStatistics
	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	return target.Items
}

func GetVideo(videoId string, apiKey string, youtubeVideoGetRequst string) []types_.ItemVideoYoutubeSnippet {
	response, err := http.Get(buildVideoYoutubeUri(videoId, "snippet", apiKey, youtubeVideoGetRequst))
	if err != nil {
		panic(err)
	}

	var target types_.TargetVideoYoutubeSnippet
	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	return target.Items
}

func buildVideoYoutubeUri(videoId string, part string, apiKey string, youtubeVideoGetRequest string) string {
	params := url.Values{
		"key":  {apiKey},
		"id":   {videoId},
		"part": {part},
	}

	return youtubeVideoGetRequest + "?" + params.Encode()
}
