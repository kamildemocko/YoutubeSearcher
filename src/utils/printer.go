package utils

import (
	"fmt"
	"strings"

	"github.com/kamildemocko/src/types_"
)

type Printer struct {
	MaxChannelLength int
	MaxTitleLength   int
}

func (p Printer) PrintHeaderLine() {
	blockChannel := strings.Repeat("#", p.MaxChannelLength+4)
	blockTitle := strings.Repeat(" ", p.MaxTitleLength)
	fmt.Printf("[ %s  %s\n\n", blockChannel, blockTitle)
}

func (p Printer) PrintItems(index int, optionData types_.CurrentData) {
	channelTitleCut := optionData.ChannelTitle
	if len(channelTitleCut) > p.MaxChannelLength {
		channelTitleCut = channelTitleCut[:p.MaxChannelLength]
	}

	videoTitleCut := optionData.Title
	if len(channelTitleCut) > p.MaxTitleLength {
		channelTitleCut = channelTitleCut[:p.MaxTitleLength]
	}

	fmt.Printf("[%02d] %*s %s\n", index+1, p.MaxChannelLength+1, channelTitleCut, videoTitleCut)
}

func (p Printer) PrintAvailableOptions() {
	fmt.Printf("\n[o OPEN VIDEO] [t OPEN THUMBNAIL] [c OPEN CHANNEL] [i SHOW DETAILS] [d DOWNLOAD VIDEO] [n NEW SEARCH] [q QUIT]\n")
}

func (p Printer) PrintInfo(title string, channel string, description string, pubishedAt string, viewCount string, likeCount string, commentCount string) {
	fmt.Printf("\nChannel name: %s\n", channel)
	fmt.Printf("Video name: %s\n", title)
	fmt.Printf("Video description: %s\n", description)
	fmt.Printf("Video published at: %s\n", pubishedAt)
	fmt.Printf("Views: %s, Likes: %s, Comments: %s\n", viewCount, likeCount, commentCount)
}
