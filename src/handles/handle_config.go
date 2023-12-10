package handles

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kamildemocko/config"
	"github.com/kamildemocko/src/utils"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	ApiKey                    string
	YoutubeSearchGetRequest   string
	YoutubeVideoGetRequtest   string
	YoutubeCommentsGetRequest string
	MaxChannelLength          int
	MaxTitleLength            int
	MaxResults                int
}

func prepareConfigPath(userDir string) string {
	var configPathPart = config.ConfigPath
	path := filepath.Join(userDir, configPathPart)
	parent := filepath.Dir(path)
	if _, err := os.Stat(parent); os.IsNotExist(err) {
		err := os.MkdirAll(parent, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return path
}

func GetConfig(userPath string) (Config, error) {
	path := prepareConfigPath(userPath)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return Config{}, fmt.Errorf("file does not exist")

	} else if err != nil {
		panic(err)
	}

	cfg := Config{}
	cfg.readConfigFromFile(path)

	return cfg, nil
}

func SaveConfig(apiKey string, youtubeSearchGetRequest string, youtubeVideoGetRequest string, youtubeCommentsGetRequest string, maxChannelLength int, maxTitleLength int, maxResults int, userDir string) {
	cfg := Config{
		ApiKey:                    apiKey,
		YoutubeSearchGetRequest:   youtubeSearchGetRequest,
		YoutubeVideoGetRequtest:   youtubeVideoGetRequest,
		YoutubeCommentsGetRequest: youtubeCommentsGetRequest,
		MaxChannelLength:          maxChannelLength,
		MaxTitleLength:            maxTitleLength,
		MaxResults:                maxResults,
	}

	path := prepareConfigPath(userDir)

	err := cfg.saveConfigToFile(path)
	if err != nil {
		panic(err)
	}
}

func (c *Config) marshalConfig(cfgStruct Config) []byte {
	b, err := toml.Marshal(cfgStruct)
	if err != nil {
		panic(err)
	}

	return b
}

func (c *Config) unmarshalConfig(b []byte) {
	err := toml.Unmarshal(b, c)
	if err != nil {
		panic(err)
	}
}

func (c *Config) saveConfigToFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	b := c.marshalConfig(*c)

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) readConfigFromFile(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	c.unmarshalConfig(b)

	return nil
}

func (c *Config) HandleFirstTime() {
	fmt.Println("Please enter configuration values, these will be saved to $HOME/.config/youtube_searcher.toml")

	apiKey, maxResults := utils.GetConfigOptions()

	*c = Config{
		ApiKey:                    apiKey,
		YoutubeSearchGetRequest:   config.DefaultYoutubeSearchGetRequest,
		YoutubeVideoGetRequtest:   config.DefaultYoutubeVideoGetRequtest,
		YoutubeCommentsGetRequest: config.DefaultYoutubeCommentsGetRequest,
		MaxChannelLength:          config.DefaultMaxChannelLength,
		MaxTitleLength:            config.DefaultMaxTitleLength,
		MaxResults:                maxResults,
	}

	SaveConfig(
		c.ApiKey,
		c.YoutubeSearchGetRequest,
		c.YoutubeVideoGetRequtest,
		c.YoutubeCommentsGetRequest,
		c.MaxChannelLength,
		c.MaxTitleLength,
		c.MaxResults,
		utils.GetUserDir(),
	)
}
