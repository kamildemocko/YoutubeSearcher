package main

import (
	"github.com/kamildemocko/src/handles"
	"github.com/kamildemocko/src/types_"
	"github.com/kamildemocko/src/utils"
)

func main() {
	config, err := handles.GetConfig(utils.GetUserDir())
	if err != nil {
		if err.Error() == "file does not exist" {
			config.HandleFirstTime()

		} else {
			panic(err)
		}
	}

	for {
		var userInput = utils.GetUserQuery()
		var isVideoId bool = handles.DetermineIsVideoId(userInput, config)

		var data []types_.CurrentData
		if isVideoId {
			data = handles.GetVideoById(userInput, config)

		} else {
			data = handles.SearchYoutube(userInput, config.MaxResults, config)
		}

		var printer = utils.Printer{
			MaxChannelLength: config.MaxChannelLength,
			MaxTitleLength:   config.MaxTitleLength,
		}

		printer.PrintHeaderLine()
		for i, v := range data {
			printer.PrintItems(i, v)
		}

		pickedOption := handles.HandleOption(&data, printer, config)
		if pickedOption == "q" {
			break
		}
	}
}
