package handler

import (
	speech "DiscordSimpleChatBot/speech"

	"github.com/bwmarrin/discordgo"
)

// MessageCreate - Detect Message Create Event
// This function will be called (due to AddHandler) every time a new
// message is created on any channel that the autenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!" {
		s.ChannelMessageSend(m.ChannelID, speech.IntroHelp)
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "!momonga" {
		s.ChannelMessageSend(m.ChannelID, speech.Introduction)
	}
}
