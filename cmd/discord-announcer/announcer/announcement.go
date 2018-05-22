package announcer

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Announcement the message we will be sending
type Announcement struct {
	Channels []string `yaml:"channels"`
	Content  string   `yaml:"content"`
	Embed    *struct {
		URL         string `yaml:"url,omitempty"`
		Type        string `yaml:"type,omitempty"`
		Title       string `yaml:"title,omitempty"`
		Description string `yaml:"description,omitempty"`
		Timestamp   string `yaml:"timestamp,omitempty"`
		Color       int    `yaml:"color,omitempty"`
		Footer      *struct {
			Text         string `yaml:"text,omitempty"`
			IconURL      string `yaml:"icon_url,omitempty"`
			ProxyIconURL string `yaml:"proxy_icon_url,omitempty"`
		}
		Image *struct {
			URL      string `yaml:"url,omitempty"`
			ProxyURL string `yaml:"proxy_url,omitempty"`
			Width    int    `yaml:"width,omitempty"`
			Height   int    `yaml:"height,omitempty"`
		}
		Thumbnail *struct {
			URL      string `yaml:"url,omitempty"`
			ProxyURL string `yaml:"proxy_url,omitempty"`
			Width    int    `yaml:"width,omitempty"`
			Height   int    `yaml:"height,omitempty"`
		}
		Video *struct {
			URL    string `yaml:"url,omitempty"`
			Width  int    `yaml:"width,omitempty"`
			Height int    `yaml:"height,omitempty"`
		}
		Provider *struct {
			URL  string `yaml:"url,omitempty"`
			Name string `yaml:"name,omitempty"`
		}
		Author *struct {
			URL          string `yaml:"url,omitempty"`
			Name         string `yaml:"name,omitempty"`
			IconURL      string `yaml:"icon_url,omitempty"`
			ProxyIconURL string `yaml:"proxy_icon_url,omitempty"`
		}
		Fields []*struct {
			Name   string `yaml:"name,omitempty"`
			Value  string `yaml:"value,omitempty"`
			Inline bool   `yaml:"inline,omitempty"`
		}
	}
}

// ParseFiles will load all the announcements
func ParseFiles(filePaths []string) ([]Announcement, error) {
	announcements := make([]Announcement, len(filePaths))

	for i, p := range filePaths {
		f, err := ioutil.ReadFile(p)
		if err != nil {
			return nil, err
		}

		var a Announcement
		if err := yaml.Unmarshal(f, &a); err != nil {
			return nil, err
		}

		announcements[i] = a
	}

	return announcements, nil
}
