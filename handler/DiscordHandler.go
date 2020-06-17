package handler

import (
	speech "DiscordSimpleChatBot/speech"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/goinvest/iexcloud/examples/iexcloud/domain"
	iex "github.com/goinvest/iexcloud/v2"
	"github.com/turnage/graw/reddit"
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

	if m.Content == "/" {
		s.ChannelMessageSend(m.ChannelID, speech.IntroHelp)
	}

	if m.Content == "/momonga" {
		s.ChannelMessageSend(m.ChannelID, speech.Introduction)
	}

	if strings.Fields(m.Content)[0] == "/meme" {
		s.ChannelMessageSend(m.ChannelID, ScrapeMemes(strings.Fields(m.Content)[1]))
	}

	if strings.Fields(m.Content)[0] == "/stock" {
		s.ChannelMessageSend(m.ChannelID, GetStockInformation(strings.Fields(m.Content)[1]))
	}
}

// ScrapeMemes - Scrape subreddit r/memes to retrieve latest memes
func ScrapeMemes(pageNumberString string) string {

	bot, err := reddit.NewBotFromAgentFile("./constants/discordbot.agent", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return err.Error()
	}

	harvest, err := bot.Listing("/r/memes", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/memes: ", err)
		return err.Error()
	}

	pageNum, err := strconv.Atoi(pageNumberString)
	if err != nil || pageNum < 0 {
		return "Page number is invalid (must be >0)"
	}

	post := harvest.Posts[pageNum]
	return post.URL
}

// GetStockInformation - Get Stock information given ticker symbol
func GetStockInformation(ticker string) string {

	cfg, err := domain.ReadConfig("./constants/iexconfig.toml")
	if err != nil {
		fmt.Println("Error reading config file: ", err)
		return err.Error()
	}

	client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
	quote, err := client.Quote(context.Background(), ticker)
	if err != nil {
		fmt.Println("Error getting quote: ", err)
		return err.Error()
	}

	return fmt.Sprint("Current price of ", quote.Symbol, " is $", quote.LatestPrice)
}
