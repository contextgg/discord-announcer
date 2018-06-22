package announcer

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/nordicgaming/discord-announcer/cmd/discord-announcer/config"
)

// Announcer for sending discord messages
type Announcer struct {
	session *discordgo.Session
}

// NewAnnouncer create a announcer
func NewAnnouncer(cfg *config.Config) (*Announcer, error) {
	var err error
	var sess *discordgo.Session

	if cfg.Discord.BotToken != "" {
		sess, err = discordgo.New("BOT " + cfg.Discord.BotToken)
		if err != nil {
			return nil, err
		}
	} else if cfg.Discord.UserToken != "" {
		sess, err = discordgo.New(cfg.Discord.UserToken)
		if err != nil {
			return nil, err
		}
	} else if cfg.Discord.Username != "" && cfg.Discord.Password != "" {
		sess, err = discordgo.New(cfg.Discord.Username, cfg.Discord.Password)
		if err != nil {
			return nil, err
		}
	}

	if sess == nil {
		return nil, errors.New("No session configured")
	}

	accouncer := &Announcer{
		session: sess,
	}

	return accouncer, nil
}

// SendAnnouncements loop through announcements and send them
func (nouncer *Announcer) SendAnnouncements(as []Announcement) error {
	for _, a := range as {
		if err := nouncer.SendAnnouncement(&a); err != nil {
			return err
		}
	}

	return nil
}

// SendAnnouncement will send an embedded message
func (nouncer *Announcer) SendAnnouncement(a *Announcement) error {
	embed := MakeEmbed(a)
	if embed == nil && a.Content == "" {
		return nil
	}

	for _, channelID := range a.Channels {
		var err error
		var msg *discordgo.Message

		send := &discordgo.MessageSend{
			Embed:   embed,
			Content: a.Content,
		}

		msg, err = nouncer.session.ChannelMessageSendComplex(channelID, send)
		if err != nil {
			return err
		}

		log.Printf("Message %s sent\n", msg.ID)
	}
	return nil
}

// MakeEmbed will turn our embed message into a discord one.
func MakeEmbed(a *Announcement) *discordgo.MessageEmbed {
	m := a.Embed
	if m == nil {
		return nil
	}

	var footer *discordgo.MessageEmbedFooter
	if m.Footer != nil {
		footer = &discordgo.MessageEmbedFooter{
			Text:         m.Footer.Text,
			IconURL:      m.Footer.IconURL,
			ProxyIconURL: m.Footer.ProxyIconURL,
		}
	}

	var image *discordgo.MessageEmbedImage
	if m.Image != nil {
		image = &discordgo.MessageEmbedImage{
			Height:   m.Image.Height,
			ProxyURL: m.Image.ProxyURL,
			URL:      m.Image.URL,
			Width:    m.Image.Width,
		}
	}

	var thumbnail *discordgo.MessageEmbedThumbnail
	if m.Thumbnail != nil {
		thumbnail = &discordgo.MessageEmbedThumbnail{
			Height:   m.Thumbnail.Height,
			ProxyURL: m.Thumbnail.ProxyURL,
			URL:      m.Thumbnail.URL,
			Width:    m.Thumbnail.Width,
		}
	}

	var video *discordgo.MessageEmbedVideo
	if m.Video != nil {
		video = &discordgo.MessageEmbedVideo{
			Height: m.Video.Height,
			URL:    m.Video.URL,
			Width:  m.Video.Width,
		}
	}

	var provider *discordgo.MessageEmbedProvider
	if m.Provider != nil {
		provider = &discordgo.MessageEmbedProvider{
			Name: m.Provider.Name,
			URL:  m.Provider.URL,
		}
	}

	var author *discordgo.MessageEmbedAuthor
	if m.Author != nil {
		author = &discordgo.MessageEmbedAuthor{
			URL:          m.Author.URL,
			Name:         m.Author.Name,
			IconURL:      m.Author.IconURL,
			ProxyIconURL: m.Author.ProxyIconURL,
		}
	}

	fields := make([]*discordgo.MessageEmbedField, len(m.Fields))
	for i, f := range m.Fields {
		fields[i] = &discordgo.MessageEmbedField{
			Name:   f.Name,
			Value:  f.Value,
			Inline: f.Inline,
		}
	}

	return &discordgo.MessageEmbed{
		URL:         m.URL,
		Type:        m.Type,
		Title:       m.Title,
		Description: m.Description,
		Timestamp:   m.Timestamp,
		Color:       m.Color,
		Footer:      footer,
		Image:       image,
		Thumbnail:   thumbnail,
		Video:       video,
		Provider:    provider,
		Author:      author,
		Fields:      fields,
	}
}
