package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func DownloadYoutubeVideo(videoId string, path string) {
	fmt.Println("Videos might not be available in their full quality")
	fmt.Println("Available videos:")

	client := youtube.Client{}
	video, err := client.GetVideo(videoId)
	if err != nil {
		fmt.Println("Error, cannot get video")
		return
	}

	chosenFormat, err := pickVideoFormatToDownload(video, "Choose number: ")
	if err != nil {
		if err.Error() == "exit" {
			return
		}

		fmt.Println("Sorry, no video format is available to download")
		return
	}

	fmt.Println("Downloading video")

	stream, _, err := client.GetStream(video, &chosenFormat)
	if err != nil {
		fmt.Println("Error downloading video, cannot get stream")
		return
	}
	defer stream.Close()

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating a file")
		return
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Println("Error saving video data to a file")
	}

	fmt.Printf("Video downloaded at %s\n", path)
}

func pickVideoFormatToDownload(v *youtube.Video, title string) (youtube.Format, error) {
	var available []youtube.Format
	var index int

	fmt.Println("0: Cancel")
	for _, f := range v.Formats.WithAudioChannels() {
		if f.QualityLabel == "" {
			continue
		}

		fmt.Printf("%d: %s\n", index+1, f.QualityLabel)
		available = append(available, f)
		index++
	}

	if index == 0 {
		return youtube.Format{}, fmt.Errorf("no video format available")
	}

	input := GetUserInputIntGeneric(1, len(available))
	if input == -1 {
		return youtube.Format{}, fmt.Errorf("exit")
	}

	return available[input-1], nil
}
