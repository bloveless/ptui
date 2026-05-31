package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/joho/godotenv"

	"github.com/bloveless/ptui/pangolin"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("main: %s", err)
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("loading .env file: %w", err)
	}
	baseURL, err := url.Parse(os.Getenv("PANGOLIN_BASE_URL"))
	if err != nil {
		return fmt.Errorf("parsing base url: %w", err)
	}
	w, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		return fmt.Errorf("opening debug log: %w", err)
	}
	logger := NewLogger(w)
	api := pangolin.NewAPI(os.Getenv("PANGOLIN_API_KEY"), baseURL)
	e := NewEventGenerator(logger, api)
	p := tea.NewProgram(initialModel(e, WithLogger(logger)))
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running tui: %w", err)
	}
	return nil
}
