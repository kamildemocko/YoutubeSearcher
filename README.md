# Youtube Searcher

## Desctiption

Simple CLI searcher for Youtube, using (own) API key.

This:

- can search by text or Youtube video ID
- can open video in a browser, thumbnail, channel (mac for now)
- show some details, like description, number of likes, comments, views
- download Youtube video, but not all quality is supported, sadly

## Why?

Of course, there is no real reason why would one use this and not visit Youtube itself, this was done mainly to learn and practise GO

## Usage

- first run will prompt you to enter youtube developer API key, this is saved to `$USER_FOLDER/.config/youtube_searcher.toml` along with few other settings
- on executing program, you will be prompted to enter search query or video ID
- list of videos has index number and program prompts you to enter an option
- by pressing option and choosing index the action is performed
- by choosing download video option, you are prompted to choose from available quality formats. These are chosen from what is available. Only formats with audio are presented. Unfortunately Higher quality formats are usually without audio.

## Screenshots
