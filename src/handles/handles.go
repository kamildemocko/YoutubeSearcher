package handles

import (
	"fmt"
	"path"
	"strings"

	"github.com/kamildemocko/src/req"
	"github.com/kamildemocko/src/types_"
	"github.com/kamildemocko/src/utils"
)

func GetVideoById(videoId string, config Config) []types_.CurrentData {
	var result = req.GetVideo(videoId, config.ApiKey, config.YoutubeVideoGetRequtest)
	var currentData = make([]types_.CurrentData, 1)

	v := result[0]
	optionData := types_.CurrentData{
		ID:           0,
		VideoId:      videoId,
		Title:        v.Snippet.Title,
		Description:  v.Snippet.Description,
		Url:          fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId),
		ThumbnailUrl: v.Snippet.Thumbnails.High.Url,
		PublishedAt:  v.Snippet.PublishedAt,
		ChannelID:    v.Snippet.ChannelId,
		ChannelTitle: v.Snippet.ChannelTitle,
		ChannelUrl:   fmt.Sprintf("https://www.youtube.com/channel/%s", v.Snippet.ChannelId),
	}

	currentData[0] = optionData

	return currentData
}

func SearchYoutube(query string, maxResults int, config Config) []types_.CurrentData {
	var querySearchResult = req.SeachYoutube(query, maxResults, config.ApiKey, config.YoutubeSearchGetRequest)
	var currentData = make([]types_.CurrentData, len(querySearchResult))

	for i, v := range querySearchResult {
		optionData := types_.CurrentData{
			ID:           i,
			VideoId:      v.Id.VideoId,
			Title:        v.Snippet.Title,
			Description:  v.Snippet.Description,
			Url:          fmt.Sprintf("https://www.youtube.com/watch?v=%s", v.Id.VideoId),
			ThumbnailUrl: v.Snippet.Thumbnails.High.Url,
			PublishedAt:  v.Snippet.PublishedAt,
			ChannelID:    v.Snippet.ChannelId,
			ChannelTitle: v.Snippet.ChannelTitle,
			ChannelUrl:   fmt.Sprintf("https://www.youtube.com/channel/%s", v.Snippet.ChannelId),
		}

		currentData[i] = optionData
	}

	return currentData
}

/*
Handles user option choice

Returns
- q for EXIT
- n for new query search
*/
func HandleOption(currentOptions *[]types_.CurrentData, printer utils.Printer, config Config) string {
	var defaultOrLast int = -1
	var option string

	for {
		printer.PrintAvailableOptions()

		option = utils.GetUserOption()
		if option == "q" || option == "n" {
			return option
		}

		userIndex := utils.GetUserInputIndex(len(*currentOptions), defaultOrLast)
		defaultOrLast = userIndex

		data := (*currentOptions)[userIndex]

		if option == "o" {
			fmt.Printf("Opening URL with video %s by %s\n", data.Title, data.ChannelTitle)
			utils.OpenLinkMac(data.Url)

		} else if option == "t" {
			fmt.Printf("Opening thumbnail of video %s by %s\n", data.Title, data.ChannelTitle)
			utils.OpenLinkMac(data.ThumbnailUrl)

		} else if option == "c" {
			fmt.Printf("Opening URL of channel %s\n", data.ChannelTitle)
			utils.OpenLinkMac(data.ChannelUrl)

		} else if option == "i" {
			var stats []types_.ItemVideoYoutubeStatistics = req.GetVideoInfo(data.VideoId, "statistics", config.ApiKey, config.YoutubeVideoGetRequtest)

			viewCount := stats[0].Statistics.ViewCount
			if viewCount == "" {
				viewCount = "0"
			}
			likeCount := stats[0].Statistics.LikeCount
			if likeCount == "" {
				likeCount = "0"
			}
			commentCount := stats[0].Statistics.CommentCount
			if commentCount == "" {
				commentCount = "0"
			}

			printer.PrintInfo(
				data.Title, data.ChannelTitle, data.Description, data.PublishedAt,
				viewCount, likeCount, commentCount,
			)

		} else if option == "d" {
			replacer := strings.NewReplacer(":", "", "\\", "")
			titleFileName := replacer.Replace(data.Title)
			path := path.Join("./downloads", titleFileName+".mp4")
			utils.DownloadYoutubeVideo(data.VideoId, path)
		}
	}
}

func DetermineIsVideoId(value string, config Config) bool {
	if len(value) != 11 {
		return false
	}

	var result []types_.ItemVideoYoutubeStatistics = req.GetVideoInfo(value, "id", config.ApiKey, config.YoutubeVideoGetRequtest)

	return len(result) == 1
}
